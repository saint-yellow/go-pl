package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saint-yellow/go-pl/toys/gohub/bootstrap"
	btsConfig "github.com/saint-yellow/go-pl/toys/gohub/config"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/config"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
    var env string
    flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
    flag.Parse()
    config.InitConfig(env)

	// 初始化 Gin 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务，指定端口为 8000
	err := router.Run(":"+config.GetString("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或其他错误
		fmt.Println(err.Error())
	}
}
