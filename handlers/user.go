package handlers

import (
	"encoding/json"

	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/slaveofcode/go-starter-api/context"
	"github.com/valyala/fasthttp"
)

// User class
type User struct {
	appCtx *context.AppContext
}

type listResponse struct {
	Items []models.User `json:"items"`
	Total int           `json:"total"`
}

// List returns list of users
func (u User) List(ctx *fasthttp.RequestCtx) {
	var entities []models.User
	var total int
	offset := 0
	limit := 20

	offsetQuery := ctx.QueryArgs().GetUintOrZero("offset")
	if offsetQuery != 0 {
		offset = offsetQuery
	}

	limitQuery := ctx.QueryArgs().GetUintOrZero("limit")
	if limitQuery != 0 {
		limit = limitQuery
	}

	u.appCtx.DB.Limit(limit).
		Offset(offset).
		Order("\"CreatedAt\" DESC").
		Find(&entities).
		Offset(0).
		Limit(-1).
		Count(&total)

	resUser, _ := json.Marshal(&listResponse{
		Items: entities,
		Total: total,
	})

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(resUser)
	return
}

// NewUser create new user instance
func NewUser(appCtx *context.AppContext) *User {
	return &User{
		appCtx: appCtx,
	}
}
