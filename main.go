package main

import (
	e "root/check_error"
	"root/db"
	r "root/routes/user"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	store, err := db.Open(db.PostgresConfiguration())
	e.ErrorHandler(err)

	store.CheckStoreConnection()

	app := router.New()

	r.User(store, app)

	fasthttp.ListenAndServe(":5000", app.Handler)
}
