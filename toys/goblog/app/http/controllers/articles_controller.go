package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/saint-yellow/go-pl/toys/goblog/pkg/logger"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/route"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/types"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

    article, err := getArticleByID(id)

    if err != nil {
        if err == sql.ErrNoRows {
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprint(w, "404 文章未找到")
        } else {
            logger.LogError(err)
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "500 服务器内部错误")
        }
    } else {
        tmpl, err := template.
            New("show.gohtml").
            Funcs(template.FuncMap{
                "RouteNameToURL": route.NameToURL,
                "Int64ToString": types.Int64ToString,
            }).
            ParseFiles("resources/views/articles/show.gohtml")
        logger.LogError(err)
        err = tmpl.Execute(w, article)
        logger.LogError(err)
}}
