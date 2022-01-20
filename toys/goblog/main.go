package main

import (
	"net/http"

	"github.com/saint-yellow/go-pl/toys/goblog/app/http/middlewares"
	"github.com/saint-yellow/go-pl/toys/goblog/bootstrap"
	"github.com/saint-yellow/go-pl/toys/goblog/pkg/logger"
)

func main() {

    bootstrap.SetupDB()
    router := bootstrap.SetupRoute()

    err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
    logger.LogError(err)
}