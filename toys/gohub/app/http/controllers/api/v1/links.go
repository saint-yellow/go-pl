package v1

import (
	"github.com/saint-yellow/go-pl/toys/gohub/app/models/link"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
    BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
    links := link.All()
    response.Data(c, links)
}
