package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/saint-yellow/go-pl/toys/memorandum/api"
	"github.com/saint-yellow/go-pl/toys/memorandum/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()


	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session0", store))

	apiGroup := router.Group("api") 
	{
		apiGroup.POST("user/register", api.RegisterUser)
		apiGroup.POST("user/login", api.LoginUser)
		
		auth := apiGroup.Group("/")
		auth.Use(middleware.ValidateJWT())
		{
			auth.GET("tasks/:pageNo/:pageSize", api.GetTasks)
			auth.GET("task/:id", api.GetTask)
			auth.POST("task", api.CreateTask)
			auth.PUT("task/:id", api.UpdateTask)
			auth.PATCH("task/:id", api.UpdateTask)
			auth.DELETE("task/:id", api.DeleteTask)
			auth.GET("tasks", api.SearchTasks)
		}
	}

	return router
}
