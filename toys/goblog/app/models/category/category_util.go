package category

import "github.com/saint-yellow/go-pl/toys/goblog/pkg/route"

// Link 方法用来生成文章链接
func (c Category) Link() string {
    return route.NameToURL("categories.show", "id", c.GetStringID())
}
