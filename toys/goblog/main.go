package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/saint-yellow/go-pl/toys/goblog/bootstrap"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/database"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/logger"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/route"
)

var ( 
    router *mux.Router
    db *sql.DB
)

type Article struct {
    Title, Body string
    ID int64
}

// Delete 方法用以从数据库中删除单条记录
func (a Article) Delete() (rowsAffected int64, err error) {
    rs, err := db.Exec("DELETE FROM articles WHERE id = " + strconv.FormatInt(a.ID, 10))

    if err != nil {
        return 0, err
    }

    // √ 删除成功，跳转到文章详情页
    if n, _ := rs.RowsAffected(); n > 0 {
        return n, nil
    }

    return 0, nil
}

func getArticleByID(id string) (Article, error) {
    article := Article{}
    query := "SELECT * FROM articles WHERE id = ?"
    err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
    return article, err
}

// 充当中间件
func forceHTMLMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 1. 设置标头
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        // 2. 继续处理请求
        next.ServeHTTP(w, r)
    })
}

func removeTailingSlash(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
        }
        next.ServeHTTP(w, r)
    })
}

func articlesDeleteHandler(w http.ResponseWriter, r *http.Request) {

    // 1. 获取 URL 参数
    // id := getRouteVariable("id", r)
    id := route.GetRouteVariable("id", r)

    // 2. 读取对应的文章数据
    article, err := getArticleByID(id)

    // 3. 如果出现错误
    if err != nil {
        if err == sql.ErrNoRows {
            // 3.1 数据未找到
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprint(w, "404 文章未找到")
        } else {
            // 3.2 数据库错误
            logger.LogError(err)
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "500 服务器内部错误")
        }
    } else {
        // 4. 未出现错误，执行删除操作
        rowsAffected, err := article.Delete()

        // 4.1 发生错误
        if err != nil {
            // 应该是 SQL 报错了
            logger.LogError(err)
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "500 服务器内部错误")
        } else {
            // 4.2 未发生错误
            if rowsAffected > 0 {
                // 重定向到文章列表页
                indexURL, _ := router.Get("articles.index").URL()
                http.Redirect(w, r, indexURL.String(), http.StatusFound)
            } else {
                // Edge case
                w.WriteHeader(http.StatusNotFound)
                fmt.Fprint(w, "404 文章未找到")
            }
        }
    }
}

func main() {
    database.Initialize()
    db = database.DB

    bootstrap.SetupDB()

    router = bootstrap.SetupRoute()

    router.HandleFunc("/articles/{id:[0-9]+}/delete", articlesDeleteHandler).Methods("POST").Name("articles.delete")

    // 中间件：强制内容类型为 HTML
    router.Use(forceHTMLMiddleware)

    http.ListenAndServe(":3000", removeTailingSlash(router))
}