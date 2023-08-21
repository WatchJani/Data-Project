package user

import (
	"root/db"

	"root/handler"

	"github.com/fasthttp/router"
)

func User(db *db.Store, app *router.Router) {
	user := handler.New(db)

	app.GET("/user", user.GetAll)
	app.GET("/user/{id}", user.GetUser)

	app.POST("/user", user.NewUser)
	// app.POST("/user/{code}", user.GetAll)

}
