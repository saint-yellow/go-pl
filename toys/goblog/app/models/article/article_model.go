package article

import "github.com/saint-yellow/go-pl/toys/goblog/app/models"

// Article 文章模型
type Article struct {
    models.BaseModel

    Title string `gorm:"type:varchar(255);not null;" valid:"title"`
    Body  string `gorm:"type:longtext;not null;" valid:"body"`
}

