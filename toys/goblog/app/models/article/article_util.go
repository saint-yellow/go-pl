package article

import (
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/route"
)

// Link 方法用来生成文章链接
func (a Article) Link() string {
    return route.NameToURL("articles.show", "id", a.GetStringID())
}