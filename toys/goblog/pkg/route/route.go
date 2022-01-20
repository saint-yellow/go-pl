package route

import "github.com/gorilla/mux"

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
}

// RouteName2URL 通过路由名称来获取 URL
func NameToURL(routeName string, pairs ...string) string {
    url, err := Router.Get(routeName).URL(pairs...)
    if err != nil {
        return ""
    }

    return url.String()
}
