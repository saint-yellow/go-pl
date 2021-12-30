package serialization

import "github.com/saint-yellow/go-pl/toys/memorandum/model"

type User struct {
	ID int
	UserName string
	Token string
}

func SerializeUser(model.User) User {
	return User{
		
	}
} 