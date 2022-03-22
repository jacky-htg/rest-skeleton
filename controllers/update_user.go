package controllers

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := new(models.User)
	user.Db = u.Db
	user.Log = u.Log
	user.ID = uint(id)

	if err = user.Get(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userRequest := new(request.UserRequest)
	if err := api.Decode(r, &userRequest); err != nil {
		u.Log.Printf("error decode user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userUpdate := userRequest.Transform(user)
	userUpdate.Db = u.Db
	userUpdate.Log = u.Log
	if err = userUpdate.Update(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := new(response.UserResponse)
	resp.Transform(*userUpdate)
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
