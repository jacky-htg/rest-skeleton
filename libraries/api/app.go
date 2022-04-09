package api

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// App is the entrypoint into our application and what controls the context of
// each request. Feel free to add any configuration data/logic on this type.
type App struct {
	log *log.Logger
	mux *httprouter.Router
	mw  []Middleware
}

// Handler type as standard http.Handle
type Handler func(http.ResponseWriter, *http.Request)

// Ctx type for encapsulated context key
type Ctx string

// Handle associates a httprouter Handle function with an HTTP Method and URL pattern.
func (a *App) Handle(method, url string, h Handler) {
	h = wrapMiddleware(a.mw, h)

	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), Ctx("ps"), ps)
		h(w, r.WithContext(ctx))
	}

	a.mux.Handle(method, url, fn)
}

// ServeHTTP implements the http.Handler interface.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

//NewApp is function to create new App
func NewApp(log *log.Logger, mw ...Middleware) *App {
	return &App{
		log: log,
		mux: httprouter.New(),
		mw:  mw,
	}
}
