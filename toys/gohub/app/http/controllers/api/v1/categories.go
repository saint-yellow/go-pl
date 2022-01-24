package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saint-yellow/go-pl/toys/gohub/app/models/category"
	"github.com/saint-yellow/go-pl/toys/gohub/app/requests"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/response"
)

type CategoriesController struct {
    BaseAPIController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

    request := requests.CategoryRequest{}
    if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
        return
    }

    categoryModel := category.Category{
        Name:        request.Name,
        Description: request.Description,
    }
    categoryModel.Create()
    if categoryModel.ID > 0 {
        response.Created(c, categoryModel)
    } else {
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}