package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/toys/memorandum/configuration"
	"github.com/saint-yellow/go-pl/toys/memorandum/router"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/saint-yellow/go-pl/toys/memorandum/docs"
)

// @title RESTful API collection for Memorandum web app
// @version 1.0
// @description This is a sample server celler server.
// @host 127.0.0.1:3000
// @BasePath /api
func main() {
	r := router.NewRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	sc := configuration.Service()
	r.Run(fmt.Sprintf("%s:%s", sc.Host, sc.Port))
}