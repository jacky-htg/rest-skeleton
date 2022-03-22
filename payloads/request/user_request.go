package request

import (
	"rest/models"
)

// NewUserRequest : format json request for new user
type NewUserRequest struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

// Transform NewUserRequest to User
func (u *NewUserRequest) Transform() *models.User {
	var user models.User
	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password

	return &user
}

type UserRequest struct {
	ID       uint `json:"id"`
	IsActive bool `json:"is_active"`
}

// Transform UserRequest to User
func (u *UserRequest) Transform(user *models.User) *models.User {
	if u.ID == user.ID {
		user.IsActive = u.IsActive
	}

	return user
}
