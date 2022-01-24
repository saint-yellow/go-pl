package user

import (
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/database"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/hash"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
    var count int64
    database.DB.Model(User{}).Where("email = ?", email).Count(&count)
    return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
    var count int64
    database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
    return count > 0
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
    return hash.BcryptCheck(_password, userModel.Password)
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(email string) (userModel User) {
    database.DB.Where("phone = ?", email).First(&userModel)
    return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
    database.DB.
        Where("phone = ?", loginID).
        Or("email = ?", loginID).
        Or("name = ?", loginID).
        First(&userModel)
    return
}
