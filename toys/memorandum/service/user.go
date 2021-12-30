package service

import (
	"fmt"
	"time"

	"github.com/saint-yellow/go-pl/toys/memorandum/middleware"
	"github.com/saint-yellow/go-pl/toys/memorandum/model"
	"github.com/saint-yellow/go-pl/toys/memorandum/serialization"
)

type RegisterUser struct {
	UserName  string `form:"userName" json:"userName" binding:"required,min=3,max=15" example:"saint-yellow"`
	Password  string `form:"password" json:"password" binding:"required,min=5,max=16" example:"saint-yellow"`
}

func (service *RegisterUser) Run() *serialization.Result {
	var user model.User
	var count int
	model.Database.Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count > 0 {
		return &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("user %s already exists", service.UserName),
		}
	}

	if err := user.SetPassword(service.Password); err != nil {
		return &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("failed to register: %s", err.Error()),
		}
	}

	user.UserName = service.UserName
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	if err := model.Database.Create(&user).Error; err != nil {
		return &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("failed to register:%s", err.Error()),
		}
	}

	return &serialization.Result{
		Message: "ok",
	}
}

type LoginUser struct {
	UserName  string `form:"userName" json:"userName" binding:"required,min=3,max=15" example:"saint-yellow"`
	Password  string `form:"password" json:"password" binding:"required,min=5,max=16" example:"saint-yellow"`
}

func (service *LoginUser) Run() *serialization.Result {
	var user model.User
	var count int
	model.Database.Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count < 1 {
		return &serialization.Result{
			Code: 1,
			Message: "user not exist",
		}
	}

	if !user.CheckPassword(service.Password) {
		return &serialization.Result{
			Code: 1,
			Message: "password incorrect",
		}
	}

	// todo: token
	token, err := middleware.GenerateJWT(user.ID, service.UserName, service.Password, 0)
	if err != nil {
		return &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("failed to generate token: %s", err.Error()),
		}
	}

	return &serialization.Result{
		Message: "ok",
		Data: token,
	}
}