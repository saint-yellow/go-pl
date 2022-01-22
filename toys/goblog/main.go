package main

import (
	"embed"
	"net/http"

	"github.com/saint-yellow/go-pl/toys/goblog/app/http/middlewares"
	"github.com/saint-yellow/go-pl/toys/goblog/bootstrap"
	"github.com/saint-yellow/go-pl/toys/goblog/config"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/logger"
)

//go:embed resources/views/articles/*
//go:embed resources/views/auth/*
//go:embed resources/views/categories/*
//go:embed resources/views/layouts/*
var tplFS embed.FS

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
    bootstrap.SetupDB()

	// 初始化模板
	bootstrap.SetupTemplate(tplFS)

	// 初始化路由绑定
    router := bootstrap.SetupRoute()

    err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
    logger.LogError(err)
}