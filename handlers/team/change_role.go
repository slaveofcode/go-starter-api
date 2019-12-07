package team

import (
	"encoding/json"

	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

type changeRoleBodyParam struct {
	TeamMemberID uint `json:"teamMemberId" validate:"required"`
	FromRoleID   uint `json:"fromRoleId" validate:"required"`
	ToRoleID     uint `json:"toRoleId" validate:"required"`
}

// ChangeMemberRole will change role of member team
func (t Team) ChangeMemberRole(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var param changeRoleBodyParam
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

		var roles []models.Role
		if err := db.Where("id IN (?)", []uint{param.FromRoleID, param.ToRoleID}).Find(&roles).Error; err != nil {
			httpresponse.JSONErr(ctx, "Invalid post data: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		var teamMember models.TeamMember
		if db.Where(&models.TeamMember{
			ID:     param.TeamMemberID,
			RoleID: param.FromRoleID,
		}).First(&teamMember).RecordNotFound() {
			httpresponse.JSONErr(ctx, "Invalid post data: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		var currentMember models.TeamMember
		if err = db.Preload("Role").Where(&models.TeamMember{
			UserID: sessionData.UserID,
		}).First(&currentMember).Error; err != nil {
			httpresponse.JSONErr(ctx, "Invalid post data: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		if currentMember.Role.Name != Owner {
			httpresponse.JSONErr(ctx, "Unauthorized member: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		if err = db.Model(&teamMember).Updates(&models.TeamMember{
			RoleID: param.ToRoleID,
		}).Error; err != nil {
			httpresponse.JSONErr(ctx, "Unable to change the role: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
	}
}
