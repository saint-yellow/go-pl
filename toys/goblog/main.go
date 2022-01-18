package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
    if r.URL.Path == "/" {
        fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
    } else {
		w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
            "<p>如有疑惑，请联系我们。</p>")
    }
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 <a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	id := strings.SplitN(r.URL.Path,"/", 3)[2]
	fmt.Fprint(w, "文章 ID:" + id)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handlerFunc)
	router.HandleFunc("/about", aboutHandler)
	router.HandleFunc("/articles/", articlesHandler)
	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            fmt.Fprintln(w, "访问文章列表")
        case "POST":
            fmt.Fprintln(w, "创建新的文章")
        }
    })

	http.ListenAndServe(":3000", router)
}
