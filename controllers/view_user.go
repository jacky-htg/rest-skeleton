package controllers

import (
	"encoding/json"
	"net/http"
	"rest/libraries/api"
	"rest/models"
	"rest/payloads/response"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// View user by id
func (u *Users) View(w http.ResponseWriter, r *http.Request) {
	paramID := r.Context().Value(api.Ctx("ps")).(httprouter.Params).ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Println("convert param to id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := new(models.User)
	user.Log = u.Log
	user.Db = u.Db
	user.ID = uint(id)
	err = user.Get()
	if err != nil {
		u.Log.Println("Get User", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := new(response.UserResponse)
	resp.Transform(*user)
	data, err := json.Marshal(resp)
	if err != nil {
		u.Log.Println("Marshall data user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}
