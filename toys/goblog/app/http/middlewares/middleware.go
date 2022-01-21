package middlewares

import "net/http"

// HttpHandlerFunc 简写 —— func(http.ResponseWriter, *http.Request)
type HTTPHandlerFunc func(http.ResponseWriter, *http.Request)
