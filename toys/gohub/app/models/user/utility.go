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
