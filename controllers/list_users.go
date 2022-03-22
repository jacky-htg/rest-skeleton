package controllers

import (
	"encoding/json"
	"net/http"
	"rest/models"
	"rest/payloads/response"
)

// ListUsers : http handler for returning list of users
func (u *Users) List(w http.ResponseWriter, r *http.Request) {

	var user models.User
	user.Db = u.Db
	user.Log = u.Log

	list, err := user.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var responseList []response.UserResponse
	for _, l := range list {
		var res response.UserResponse
		res.Transform(l)
		responseList = append(responseList, res)
	}

	data, err := json.Marshal(responseList)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}
