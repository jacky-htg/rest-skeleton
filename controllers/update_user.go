package controllers

import (
	"net/http"
	"rest/libraries/api"
	"rest/models"
	"rest/payloads/request"
	"rest/payloads/response"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	paramID := r.Context().Value(api.Ctx("ps")).(httprouter.Params).ByName("id")
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

	if err = user.Get(); err != nil {
		api.ResponseError(w, err)
		return
	}

	userRequest := new(request.UserRequest)
	if err := api.Decode(r, &userRequest); err != nil {
		u.Log.Printf("error decode user: %s", err)
		api.ResponseError(w, err)
		return
	}

	userUpdate := userRequest.Transform(user)
	userUpdate.Db = u.Db
	userUpdate.Log = u.Log
	if err = userUpdate.Update(); err != nil {
		api.ResponseError(w, err)
		return
	}

	resp := new(response.UserResponse)
	resp.Transform(*userUpdate)
	if err := api.ResponseOK(w, resp, http.StatusOK); err != nil {
		u.Log.Println(err)
		api.ResponseError(w, err)
		return
	}
}
