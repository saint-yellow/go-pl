// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/saint-yellow/go-pl/toys/gohub/app/http/controllers/api/v1"
	"github.com/saint-yellow/go-pl/toys/gohub/app/models/user"
	"github.com/saint-yellow/go-pl/toys/gohub/app/requests"
)

// SignupController 注册控制器
type SignupController struct {
    v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
    // 获取请求对象，并做表单验证
    request := requests.SignupPhoneExistRequest{}
    if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
        return
    }

    // 检查数据库并返回响应
    c.JSON(http.StatusOK, gin.H{
        "exist": user.IsPhoneExist(request.Phone),
    })
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
    // 获取请求对象，并做表单验证
    request := requests.SignupEmailExistRequest{}
    if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
        return
    }

    // 检查数据库并返回响应
    c.JSON(http.StatusOK, gin.H{
        "exist": user.IsEmailExist(request.Email),
    }) 
}