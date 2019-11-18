package team

import (
	"encoding/json"

	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

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

		var team models.Team
		if db.Where(&models.TeamMember{
			UserID: sessionData.UserID,
			RoleID: roleOwner.ID,
		}).First(&team).RecordNotFound() {
			tx := t.appCtx.DB.Begin()
			// create team
			createdTeam := new(models.Team)
			if err := tx.Create(&models.Team{
				Name: param.Name,
			}).Scan(&createdTeam).Error; err != nil {
				tx.Rollback()
				httpresponse.JSONErr(ctx, "Unable to create team: "+err.Error(), fasthttp.StatusInternalServerError)
				return
			}

			// create team member as owner
			if err = tx.Create(&models.TeamMember{
				TeamID: createdTeam.ID,
				UserID: sessionData.UserID,
				RoleID: roleOwner.ID,
			}).Error; err != nil {
				tx.Rollback()
				httpresponse.JSONErr(ctx, "Unable to create team: "+err.Error(), fasthttp.StatusInternalServerError)
				return
			}

			tx.Commit()
			httpresponse.JSONOk(ctx, fasthttp.StatusCreated)
			return
		}

		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
		return
	}
}
