package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/saint-yellow/go-pl/toys/gohub/app/models/user"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/logger"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	_user := user.GetByMulti(email)
	if _user.ID == 0 {
		return user.User{}, errors.New("帐号不存在")
	}

	if !_user.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return _user, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	_user := user.GetByPhone(phone)
	if _user.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return _user, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
    userModel, ok := c.MustGet("current_user").(user.User)
    if !ok {
        logger.LogIf(errors.New("无法获取用户"))
        return user.User{}
    }
    // db is now a *DB value
    return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
    return c.GetString("current_user_id")
}
