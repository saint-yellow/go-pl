package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saint-yellow/go-pl/toys/memorandum/serialization"
	"github.com/saint-yellow/go-pl/toys/memorandum/service"
)

// @Summary 用户注册
// @Description 用户注册
// @Accept json
// @Param data body service.RegisterUser true "用户名, 密码"
// @Success 200 {object} serialization.Result "{"message":"ok"}"
// @Failure 400 {object} serialization.Result "{"message":"err","code":1}"
// @Router /api/user/register [post]
func RegisterUser(c *gin.Context) {
	srv := service.RegisterUser{}
	err := c.ShouldBind(&srv)
	if err != nil {
		c.JSON(200, &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("parameter error: %s", err.Error()),
		})
	}

	res := srv.Run()
	c.JSON(200, res)
}

// @Summary 用户登录
// @Description 用户登录
// @Accept json
// @Param data body service.LoginUser true "用户名, 密码"
// @Success 200 {object} serialization.Result "{"message":"ok"}"
// @Failure 400 {object} serialization.Result "{"message":"err","code":1}"
// @Router /api/user/login [post]
func LoginUser(c * gin.Context) {
	srv := service.LoginUser{}
	err := c.ShouldBind(&srv)
	if err != nil {
		c.JSON(200, &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("parameter error: %s", err.Error()),
		})
	}

	res := srv.Run()
	c.JSON(200, res)
}