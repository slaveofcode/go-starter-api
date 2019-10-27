package user

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
)

// List returns list of users
func (u User) List(ctx *fasthttp.RequestCtx) {
	var entities []models.User
	var total int
	offset := 0
	var limit int64 = 20

	offsetQuery := ctx.QueryArgs().GetUintOrZero("offset")
	if offsetQuery != 0 {
		offset = offsetQuery
	}

	limitQuery := ctx.QueryArgs().GetUintOrZero("limit")
	if limitQuery != 0 {
		limit = int64(limitQuery)
	}

	store, err := u.appCtx.Sesssion.Get(ctx)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	defer u.appCtx.Sesssion.Save(ctx, store)

	u.appCtx.DB.Limit(limit).
		Offset(offset).
		Order("\"CreatedAt\" DESC").
		Find(&entities).
		Offset(0).
		Limit(-1).
		Count(&total)

	httpresponse.JSON(ctx, &listResponse{
		Items: entities,
		Total: total,
	}, fasthttp.StatusOK)
	return
}
