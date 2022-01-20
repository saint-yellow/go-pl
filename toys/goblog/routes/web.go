package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saint-yellow/go-pl/toys/goblog/app/http/controllers"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

    // 静态页面
	pc := new(controllers.PagesController)
    r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
    r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
    r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
}