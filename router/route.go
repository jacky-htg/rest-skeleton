package router

import (
	"database/sql"
	"log"
	"net/http"
	"rest/controllers"
	"rest/libraries/api"
	"rest/middlewares"
)

func mid(db *sql.DB, log *log.Logger) []api.Middleware {
	var mw []api.Middleware
	mw = append(mw, middlewares.Auths(db, log, []string{"/login"}))

	return mw
}

// API : implement a http.Handler interface
func API(db *sql.DB, log *log.Logger) http.Handler {
	app := api.NewApp(log, mid(db, log)...)
	users := controllers.Users{Db: db, Log: log}

	app.Handle(http.MethodGet, "/users", users.List)
	app.Handle(http.MethodPost, "/users", users.Create)
	app.Handle(http.MethodGet, "/users/:id", users.View)
	app.Handle(http.MethodPut, "/users/:id", users.Update)
	app.Handle(http.MethodDelete, "/users/:id", users.Delete)

	return app
}
