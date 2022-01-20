package route

import "github.com/gorilla/mux"

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
}

// RouteNameToURL 通过路由名称来获取 URL
func RouteNameToURL(routeName string, pairs ...string) string {
    url, err := Router.Get(routeName).URL(pairs...)
    if err != nil {
        // checkError(err)
        return ""
    }

    return url.String()
}