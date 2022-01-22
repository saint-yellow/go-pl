package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/config"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/logger"
)

var Router *mux.Router

func SetRouter(r *mux.Router) {
    Router = r
}

// RouteNameToURL 通过路由名称来获取 URL
func NameToURL(routeName string, pairs ...string) string {
    url, err := Router.Get(routeName).URL(pairs...)
    if err != nil {
        logger.LogError(err)
        return ""
    }

    return config.GetString("app.url") + url.String()
}

// GetRouteVariable 获取 URI 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
    vars := mux.Vars(r)
    return vars[parameterName]
}

