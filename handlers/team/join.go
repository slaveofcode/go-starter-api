package team

import (
	"time"
	"encoding/json"

	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

type joinTeamBodyParam struct {
	Token string `json:"token" validate:"required"`
}

// JoinTeam will add invited member to the team
func (t Team) JoinTeam(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var param joinTeamBodyParam
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

		var invitation models.TeamMemberInvitation
		if db.Where(&models.TeamMemberInvitation{
			InvitationKey: param.Token,
		}).First(&invitation).RecordNotFound() {
			httpresponse.JSONErr(ctx, "Invalid Request", fasthttp.StatusBadRequest)
			return
		}

		if invitation.VisitedAt != nil {
			httpresponse.JSONErr(ctx, "Invalid Token", fasthttp.StatusBadRequest)
			return
		}

		var cred models.Credential
		if db.Preload("User").Where(&models.Credential{
			Email: invitation.Email,
		}).First(&cred).RecordNotFound() {
			httpresponse.JSONErr(ctx, "Please create an account to get joined", fasthttp.StatusBadRequest)
			return
		}

		var member models.TeamMember
		if db.Where(&models.TeamMember{
			TeamID: invitation.TeamID,
			UserID: cred.User.ID,
			RoleID: invitation.RoleID,
		}).First(&member).RecordNotFound() {
			// create team member
			tx := db.Begin()

			if err := tx.Create(&models.TeamMember{
				TeamID: invitation.TeamID,
				UserID: cred.User.ID,
				RoleID: invitation.RoleID,
			}).Error; err != nil {
				tx.Rollback()
				httpresponse.JSONErr(ctx, "Unable to process join: "+err.Error(), fasthttp.StatusInternalServerError)
				return
			}

			// mark invitation as visited
			now := time.Now()
			if err = tx.Model(&invitation).Updates(models.TeamMemberInvitation{
				VisitedAt: &now,
			}).Error; err != nil {
				tx.Rollback()
				httpresponse.JSONErr(ctx, "Unable to process join: "+err.Error(), fasthttp.StatusInternalServerError)
				return
			}

			tx.Commit()

			httpresponse.JSONOk(ctx, fasthttp.StatusOK)
			return
		}

		httpresponse.JSONErr(ctx, "Already joined to the team", fasthttp.StatusBadRequest)
		return
	}
}
