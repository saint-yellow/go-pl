package user

// ComparePassword 对比密码是否匹配
func (user *User) ComparePassword(password string) bool {
    return user.Password == password
}
