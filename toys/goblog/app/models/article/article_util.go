package article

import (
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/route"
)

// Link 方法用来生成文章链接
func (a Article) Link() string {
    return route.NameToURL("articles.show", "id", a.GetStringID())
}

// CreatedAtDate 创建日期
func (article Article) CreatedAtDate() string {
    return article.CreatedAt.Format("2006-01-02")
}
