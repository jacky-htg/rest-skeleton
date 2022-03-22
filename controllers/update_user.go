package controllers

import (
	"net/http"
	"rest/libraries/api"

	"github.com/julienschmidt/httprouter"
)

func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(api.Ctx("ps")).(httprouter.Params).ByName("id")
	println(id)
}
