package policies

import (
	"github.com/saint-yellow/go-pl/toys/goblog/app/models/article"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/auth"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article article.Article) bool {
    return auth.User().ID == _article.UserID
}
