// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saint-yellow/go-pl/toys/gohub/app/http/controllers/api/v1/auth"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

    // 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
    v1 := r.Group("/v1")
    {
        authGroup := v1.Group("/auth")
        {
            // 注册相关的接口
            suc := new(auth.SignupController)
            // 判断手机是否已经注册
            authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
            // 判断邮箱是否已经注册
            authGroup.POST("/signup/email/exist", suc.IsEmailExist)
            // 使用手机注册
            authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
            // 使用邮箱注册
            authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

            // 验证码相关的接口
            vcc := new(auth.VerifyCodeController)
            // 图片验证码，需要加限流
            authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
            // 发送短信验证码
            authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
            // 发送邮件验证码
            authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
        }
    }
}
