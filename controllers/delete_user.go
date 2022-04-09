package controllers

import (
	"net/http"
	"rest/libraries/api"
	"rest/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Delete user by id
func (u *Users) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	paramID := ctx.Value(api.Ctx("ps")).(httprouter.Params).ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Println("convert param to id", err)
		api.ResponseError(w, err)
		return
	}

	user := new(models.User)
	user.Db = u.Db
	user.Log = u.Log
	user.ID = uint(id)
	err = user.Get(ctx)
	if err != nil {
		api.ResponseError(w, err)
		return
	}

	err = user.Delete(ctx)
	if err != nil {
		api.ResponseError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.WriteHeader(http.StatusNoContent)
	api.ResponseOK(w, nil, http.StatusOK)
}
