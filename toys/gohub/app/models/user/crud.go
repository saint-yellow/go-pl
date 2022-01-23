package user

import "github.com/saint-yellow/go-pl/toys/gohub/pkg/database"

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
    database.DB.Create(&userModel)
}
