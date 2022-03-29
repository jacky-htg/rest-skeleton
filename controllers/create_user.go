package controllers

import (
	"fmt"
	"net/http"
	"rest/libraries/api"
	"rest/payloads/request"
	"rest/payloads/response"

	"golang.org/x/crypto/bcrypt"
)

// Create new user
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.NewUserRequest

	if err := api.Decode(r, &userRequest); err != nil {
		u.Log.Printf("error decode user: %s", err)
		api.ResponseError(w, err)
		return
	}

	if userRequest.Password != userRequest.RePassword {
		err := api.ErrBadRequest(fmt.Errorf("Password not match"), "")
		u.Log.Printf("error : %s", err)
		api.ResponseError(w, err)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Printf("error generate password: %s", err)
		api.ResponseError(w, err)
		return
	}

	userRequest.Password = string(pass)

	user := userRequest.Transform()
	user.Db = u.Db
	user.Log = u.Log

	if err := user.Create(); err != nil {
		u.Log.Printf("error call create user: %s", err)
		api.ResponseError(w, err)
		return
	}

	var res response.UserResponse
	res.Transform(*user)
	if err := api.ResponseOK(w, res, http.StatusCreated); err != nil {
		u.Log.Println(err)
		api.ResponseError(w, err)
		return
	}
}
