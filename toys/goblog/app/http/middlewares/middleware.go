package middlewares

import "net/http"

// HttpHandlerFunc įŽå ââ func(http.ResponseWriter, *http.Request)
type HTTPHandlerFunc func(http.ResponseWriter, *http.Request)
