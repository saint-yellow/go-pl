package middlewares

import (
	"net/http"

	"github.com/saint-yellow/go-pl/toys/goblog/pkg/session"
)

func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 启动会话
		session.StartSession(w, r)

		// 2. 继续处理接下去的请求
		next.ServeHTTP(w, r)
	})
}
