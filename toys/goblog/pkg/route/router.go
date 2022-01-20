package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

// RouteName2URL 通过路由名称来获取 URL
func NameToURL(routeName string, pairs ...string) string {
    url, err := Router.Get(routeName).URL(pairs...)
    if err != nil {
        return ""
    }

    return url.String()
}

// GetRouteVariable 获取 URI 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
    vars := mux.Vars(r)
    return vars[parameterName]
}

