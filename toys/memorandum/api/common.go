package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saint-yellow/go-pl/toys/memorandum/service"
)

func runService(c *gin.Context, s service.Service) {
	res := s.Run()
	c.JSON(http.StatusOK, res)
}
