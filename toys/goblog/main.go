package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/saint-yellow/go-pl/toys/goblog/bootstrap"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/database"
)

var ( 
    router *mux.Router
)

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

func main() {
    database.Initialize()

    bootstrap.SetupDB()

    router = bootstrap.SetupRoute()

    // 中间件：强制内容类型为 HTML
    router.Use(forceHTMLMiddleware)

    http.ListenAndServe(":3000", removeTailingSlash(router))
}