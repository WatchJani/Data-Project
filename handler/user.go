package handler

import (
	"root/db"
	model "root/model/user"

	e "root/check_error"

	"github.com/valyala/fasthttp"
)

type UserDB struct {
	*db.Store
}

func New(db *db.Store) *UserDB {
	return &UserDB{
		db,
	}
}

func (db *UserDB) GetAll(ctx *fasthttp.RequestCtx) {
	user, err := model.GetAll(db.DB)

	e.ErrorHandler(err)

	ctx.SetContentType("application/json")
	ctx.SetBody(user)
}

func (db *UserDB) GetUser(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue("id").(string)

	user, err := model.GetUser(db.DB, userID)
	e.ErrorHandler(err)

	ctx.SetContentType("application/json")
	ctx.SetBody(user)
}
