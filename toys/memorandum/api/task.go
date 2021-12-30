package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saint-yellow/go-pl/toys/memorandum/middleware"
	"github.com/saint-yellow/go-pl/toys/memorandum/serialization"
	"github.com/saint-yellow/go-pl/toys/memorandum/service"
)

func GetTasks(c *gin.Context) {
	claims, err := middleware.ParseJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusOK, &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("authorization error: %s", err.Error()),
		})
		return
	}

	srv := service.GetTasks{}
	srv.PageNo, _ = strconv.Atoi(c.Param("pageNo"))
	srv.PageSize, _ = strconv.Atoi(c.Param("pageSize"))
	srv.UserID = int(claims.UserID)
	
	runService(c, &srv)
}

func GetTask(c *gin.Context) {
	claims, err := middleware.ParseJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusOK, serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("authorization error:%s", err.Error()),
		})
		return
	}

	srv := service.GetTask{}
	srv.TaskID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("parameter error: %s", err.Error()),
		})
		return
	}
	srv.UserID = int(claims.UserID)
	
	runService(c, &srv)
}

func CreateTask(c *gin.Context) {
	claims, err := middleware.ParseJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusOK, serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("authorization error: %s", err.Error()),
		})
		return
	}

	srv := service.CreateTask{}
	err = c.ShouldBind(&srv)
	if err != nil {
		c.JSON(http.StatusOK, serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("parameter error: %s", err.Error()),
		})
		return
	}

	srv.UserID = int(claims.UserID)
	
	runService(c, &srv)
}

func UpdateTask(c *gin.Context) {
	claims, err := middleware.ParseJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusOK, serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("authorization error: %s", err.Error()),
		})
		return
	}

	srv := service.UpdateTask{}
	err = c.ShouldBind(&srv)
	if err != nil {
		c.JSON(http.StatusOK, serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("parameter error: %s; parameter: %v", err.Error(), srv),
		})
		return
	}

	srv.TaskID, _ = strconv.Atoi(c.Param("id"))
	srv.UserID = int(claims.UserID)

	runService(c, &srv)
}

func DeleteTask(c *gin.Context) {
	claims, err := middleware.ParseJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusOK, serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("authorization error: %s", err.Error()),
		})
		return
	}

	srv := service.DeleteTask{
		UserID: int(claims.UserID),
	}
	srv.TaskID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("parameter error: %s; parameter: %v", err.Error(), srv),
		})
		return
	}

	runService(c, &srv)
}

func SearchTasks(c *gin.Context) {
	claims, err := middleware.ParseJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusOK, serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("authorization error: %s", err.Error()),
		})
		return
	}

	srv := service.SearchTask{
		UserID: int(claims.UserID),
		Keyword: c.Query("keyword"),
	}
	if srv.PageNo, err = strconv.Atoi(c.Query("pageNo")); err != nil {
		srv.PageNo = 1
	}
	if srv.PageSize, err = strconv.Atoi(c.Query("pageSize")); err != nil {
		srv.PageSize = 20
	}

	runService(c, &srv)
}