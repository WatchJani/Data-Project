package user

import (
	"root/db"

	"root/handler"

	"github.com/fasthttp/router"
)

func User(db *db.Store, app *router.Router) {
	user := handler.New(db)

	app.GET("/user", user.GetAll)
	// app.GET("/user/{id}", user.GetAll)

	// app.POST("/user/create", user.GetAll)
	// app.POST("/user/{code}", user.GetAll)

}
