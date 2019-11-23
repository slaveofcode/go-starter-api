package team

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

const allowHaveManyTeam = false

type createTeamBodyParam struct {
	Name string `json:"name" validate:"required"`
}

// CreateTeam will create new team if not exist
func (t Team) CreateTeam(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var param createTeamBodyParam
		err := json.Unmarshal(ctx.PostBody(), &param)
		if err != nil {
			httpresponse.JSONErr(ctx, "Wrong post data: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		v := validator.New()
		err = v.Struct(param)
		if err != nil {
			httpresponse.JSONErr(ctx, "Invalid post data: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		db := t.appCtx.DB

		// create team if not exist on the user as an owner
		var roleOwner models.Role
		if err := db.Where(&models.Role{
			Name: Owner,
		}).First(&roleOwner).Error; err != nil {
			httpresponse.JSONErr(ctx, "Unable to get the role: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		var teamMember models.TeamMember
		if db.Where(&models.TeamMember{
			UserID: sessionData.UserID,
			RoleID: roleOwner.ID,
		}).First(&teamMember).RecordNotFound() {
			tx := t.appCtx.DB.Begin()
			err = createTeam(tx, param.Name, sessionData.UserID, roleOwner.ID)

			if err != nil {
				httpresponse.JSONErr(ctx, "Unable to create team: "+err.Error(), fasthttp.StatusInternalServerError)
				return
			}

			httpresponse.JSONOk(ctx, fasthttp.StatusCreated)
			return
		}

		// searching on team members already created
		var allTeamMember []models.TeamMember
		if err := db.Where(&models.TeamMember{
			UserID: sessionData.UserID,
			RoleID: roleOwner.ID,
		}).Find(&allTeamMember).Error; err != nil {
			httpresponse.JSONErr(ctx, "Unable to get member info: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		var teamIds []uint
		for _, member := range allTeamMember {
			teamIds = append(teamIds, member.TeamID)
		}

		var teams []models.Team
		if err = db.Where(teamIds).Find(&teams).Error; err != nil {
			httpresponse.JSONErr(ctx, "Unable to get teams info: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		for _, team := range teams {
			if team.Name == param.Name {
				httpresponse.JSONErr(ctx, "Already created team", fasthttp.StatusBadRequest)
				return
			}
		}

		if allowHaveManyTeam {
			tx := t.appCtx.DB.Begin()
			err = createTeam(tx, param.Name, sessionData.UserID, roleOwner.ID)

			if err != nil {
				httpresponse.JSONErr(ctx, "Unable to create team: "+err.Error(), fasthttp.StatusInternalServerError)
				return
			}

			httpresponse.JSONOk(ctx, fasthttp.StatusCreated)
			return
		}

		httpresponse.JSONErr(ctx, "You're not allowed to make more than one team", fasthttp.StatusBadRequest)
		return
	}
}

func createTeam(tx *gorm.DB, teamName string, userID, roleOwnerID uint) error {
	// create team
	createdTeam := new(models.Team)
	if err := tx.Create(&models.Team{
		Name: teamName,
	}).Scan(&createdTeam).Error; err != nil {
		tx.Rollback()
		return err
	}

	// create team member as owner
	if err := tx.Create(&models.TeamMember{
		TeamID: createdTeam.ID,
		UserID: userID,
		RoleID: roleOwnerID,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
