package user

import (
	"root/db"

	"root/handler"

	"github.com/fasthttp/router"
)

func User(db *db.Store, app *router.Router) {
	user := handler.New(db)

	app.GET("/", user.GetAll)
}
