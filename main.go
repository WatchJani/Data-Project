package main

import (
	e "root/check_error"
	"root/db"
	r "root/routes/user"

	h "root/helper"

	_ "root/sql"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	store, err := db.Open(db.PostgresConfiguration())
	e.ErrorHandler(err)

	store.CheckStoreConnection()

	defer store.Close()

	app := router.New()

	r.User(store, app)

	fasthttp.ListenAndServe(h.Port(), app.Handler)
}
