package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saint-yellow/go-pl/toys/gohub/app/http/controllers/api/v1"
	"github.com/saint-yellow/go-pl/toys/gohub/app/requests"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/auth"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/jwt"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/response"
)

// LoginController 用户控制器
type LoginController struct {
    v1.BaseAPIController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {

    // 1. 验证表单
    request := requests.LoginByPhoneRequest{}
    if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
        return
    }

    // 2. 尝试登录
    user, err := auth.LoginByPhone(request.Phone)
    if err != nil {
        // 失败，显示错误提示
        response.Error(c, err, "账号不存在或密码错误")
    } else {
        // 登录成功
        token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

        response.JSON(c, gin.H{
            "token": token,
        })
    }
}
