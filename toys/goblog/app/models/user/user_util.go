package user

import "github.com/saint-yellow/go-pl/toys/goblog/pkg/password"

// ComparePassword 对比密码是否匹配
func (user *User) ComparePassword(_password string) bool {
    return password.CheckHash(_password, user.Password)
}
