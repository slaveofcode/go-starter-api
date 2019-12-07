package team

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
)

type member struct {
	UserName string `json:"name"`
	RoleName string `json:"role"`
}

type team struct {
	Name        string   `json:"teamName"`
	TeamMembers []member `json:"members"`
}

// Members return list member
func (t Team) Members(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		db := t.appCtx.DB
		// get all team members
		var teamMembers []models.TeamMember
		if err := db.Preload("User").Preload("Role").Where(&models.TeamMember{
			UserID: sessionData.UserID,
		}).Find(&teamMembers).Error; err != nil {
			httpresponse.JSONErr(ctx, "Unable to get teams info: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		var teamIds []uint
		for _, m := range teamMembers {
			teamIds = append(teamIds, m.TeamID)
		}

		var teams []models.Team
		if err := db.Where("id IN (?)", teamIds).Find(&teams).Error; err != nil {
			httpresponse.JSONErr(ctx, "Unable to get teams info: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		var teamResult []team
		for _, t := range teams {
			team := team{
				Name: t.Name,
			}

			var members []member
			for _, m := range teamMembers {
				if m.TeamID == t.ID {
					members = append(members, member{
						UserName: m.User.Name,
						RoleName: m.Role.Name,
					})
				}
			}

			team.TeamMembers = members

			teamResult = append(teamResult, team)
		}

		httpresponse.JSON(ctx, &teamResult, fasthttp.StatusOK)
		return
	}
}
