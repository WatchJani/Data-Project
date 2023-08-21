package handler

import (
	"encoding/json"
	"root/db"
	model "root/model/user"
	"time"

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

func (db *UserDB) NewUser(ctx *fasthttp.RequestCtx) {
	requestBody := ctx.PostBody()

	request := struct {
		Name     string    `json:"Name"`
		LastName string    `json:"LastName"`
		Password string    `json:"Password"`
		Phone    string    `json:"Phone"`
		Mail     string    `json:"Mail"`
		UserName string    `json:"UserName"`
		BornDate time.Time `json:"BornDate"`
	}{}

	if e.ParserError(json.Unmarshal(requestBody, &request), ctx) {
		return
	}

	model.NewUser(db.DB, request.Name, request.LastName, request.Password, request.Phone, request.Mail, request.UserName, request.BornDate)

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBodyString(string(requestBody))
}
