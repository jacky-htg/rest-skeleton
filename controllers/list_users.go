package controllers

import (
	"net/http"
	"rest/libraries/api"
	"rest/models"
	"rest/payloads/response"
)

// ListUsers : http handler for returning list of users
func (u *Users) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user models.User
	user.Db = u.Db
	user.Log = u.Log

	list, err := user.List(ctx)
	if err != nil {
		api.ResponseError(w, err)
		return
	}

	var responseList []response.UserResponse
	for _, l := range list {
		var res response.UserResponse
		res.Transform(l)
		responseList = append(responseList, res)
	}

	if err := api.ResponseOK(w, responseList, http.StatusOK); err != nil {
		u.Log.Println(err)
		api.ResponseError(w, err)
		return
	}
}
