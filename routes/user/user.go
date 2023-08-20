package user

import (
	"root/db"

	"github.com/fasthttp/router"
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

func User(db *db.Store, app *router.Router) {
	user := New(db)

	app.GET("/", user.GetAll)
}

func (db *UserDB) GetAll(ctx *fasthttp.RequestCtx) {

}
