package controllers

import (
	"net/http"
	"rest/libraries/api"
	"rest/models"
	"rest/payloads/response"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// View user by id
func (u *Users) View(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	paramID := ctx.Value(api.Ctx("ps")).(httprouter.Params).ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Println("convert param to id", err)
		api.ResponseError(w, err)
		return
	}

	user := new(models.User)
	user.Log = u.Log
	user.Db = u.Db
	user.ID = uint(id)
	err = user.Get(ctx)
	if err != nil {
		u.Log.Println("Get User", err)
		api.ResponseError(w, err)
		return
	}

	resp := new(response.UserResponse)
	resp.Transform(*user)
	if err := api.ResponseOK(w, resp, http.StatusOK); err != nil {
		u.Log.Println(err)
		api.ResponseError(w, err)
		return
	}
}
