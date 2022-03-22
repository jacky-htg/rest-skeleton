package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest/payloads/request"
	"rest/payloads/response"

	"golang.org/x/crypto/bcrypt"
)

// Create new user
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.NewUserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userRequest)
	if err != nil {
		u.Log.Printf("error decode user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userRequest.Password != userRequest.RePassword {
		err = fmt.Errorf("Password not match")
		u.Log.Printf("error : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Printf("error generate password: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userRequest.Password = string(pass)

	user := userRequest.Transform()
	user.Db = u.Db
	user.Log = u.Log

	err = user.Create()
	if err != nil {
		u.Log.Printf("error call create user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var res response.UserResponse
	res.Transform(*user)
	data, err := json.Marshal(res)
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
