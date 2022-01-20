package article

import "github.com/saint-yellow/go-pl/toys/goblog/app/models"

// Article 文章模型
type Article struct {
    models.BaseModel
    // ID    uint64 
    Title string
    Body  string
}

