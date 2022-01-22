package article

import (
	"github.com/saint-yellow/go-pl/toys/goblog/app/models"
	"github.com/saint-yellow/go-pl/toys/goblog/app/models/user"
)

// Article 文章模型
type Article struct {
    models.BaseModel

    Title string `gorm:"type:varchar(255);not null;" valid:"title"`
    Body  string `gorm:"type:longtext;not null;" valid:"body"`

    UserID uint64 `gorm:"not null;index"`
    User user.User
}
