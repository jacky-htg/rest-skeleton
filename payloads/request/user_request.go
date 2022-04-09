package request

import (
	"errors"
	"rest/libraries/api"
	"rest/models"

	"gopkg.in/go-playground/validator.v9"
)

// NewUserRequest : format json request for new user
type NewUserRequest struct {
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"re_password" validate:"required"`
}

// Transform NewUserRequest to User
func (u *NewUserRequest) Transform() *models.User {
	var user models.User
	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password

	return &user
}

// Validate NewUserRequest
func (u *NewUserRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, verr := range err.(validator.ValidationErrors) {
			err = errors.New(verr.Field() + " is " + verr.Tag())
			break
		}

		if err != nil {
			return api.ErrBadRequest(err, err.Error())
		}
	}

	return nil
}

type UserRequest struct {
	ID       uint `json:"id" validate:"required"`
	IsActive bool `json:"is_active" validate:"required"`
}

// Transform UserRequest to User
func (u *UserRequest) Transform(user *models.User) *models.User {
	if u.ID == user.ID {
		user.IsActive = u.IsActive
	}

	return user
}
