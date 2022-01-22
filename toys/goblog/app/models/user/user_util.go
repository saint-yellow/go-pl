package user

import (
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/password"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/route"
)

// ComparePassword 对比密码是否匹配
func (user *User) ComparePassword(_password string) bool {
    return password.CheckHash(_password, user.Password)
}

// Link 方法用来生成用户链接
func (user User) Link() string {
    return route.NameToURL("users.show", "id", user.GetStringID())
}
