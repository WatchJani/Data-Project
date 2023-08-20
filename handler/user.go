package handler

import (
	"root/db"

	"github.com/valyala/fasthttp"
)

type UserDB struct {
	db *db.Store
}

func New(db *db.Store) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (db *UserDB) GetAll(ctx *fasthttp.RequestCtx) {

}
