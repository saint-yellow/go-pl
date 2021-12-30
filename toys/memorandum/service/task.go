package service

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/saint-yellow/go-pl/toys/memorandum/model"
	"github.com/saint-yellow/go-pl/toys/memorandum/serialization"
)

type GetTasks struct {
	UserID int
	PageNo int `form:"pageNo" json:"pageNo" binding:"required"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"`
}

func (s *GetTasks) Run() *serialization.Result {
	var tasks []model.Task
	total := 0

	model.Database.Model(model.Task{}).Where("user_id=?", s.UserID).Count(&total).Limit(s.PageSize).Offset((s.PageNo - 1) * s.PageSize).Find(&tasks)
	return &serialization.Result{
		Code: 0,
		Message: "ok",
		Data: serialization.SerializeTasks(tasks),
		Page: s.PageNo,
		Total: total,
	}
}

type GetTask struct {
	UserID int
	TaskID int `json:"taskId" binding:"required"` 
}

func (s *GetTask) Run() *serialization.Result {
	var task model.Task
	count := 0
	model.Database.Model(model.Task{}).Where("id=?", s.TaskID).Where("user_id=?", s.UserID).Count(&count).Find(&task)

	if count < 1 {
		return &serialization.Result{
			Code: 1,
			Message: "task not found",
		}
	}

	task.AddBrowseCount()

	return &serialization.Result{
		Code: 0,
		Message: "ok",
		Data: serialization.SerializeTask(task),
	}
}

type CreateTask struct {
	UserID int
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (s *CreateTask) Run() *serialization.Result {
	task := model.Task{
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
		UserID: s.UserID,
		Title: s.Title,
		Content: s.Content,
		Status: 0,
		StartTime: time.Now().Unix(),
	}
	err := model.Database.Create(&task).Error
	if err != nil {
		return &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("failed to create task: %s", err.Error()),
		}
	}
	return &serialization.Result{
		Code: 0,
		Message: "ok",
		Data: serialization.SerializeTask(task),
	}
}


type UpdateTask struct {
	UserID int
	TaskID int `form:"taskId" json:"taskId"`
	Title string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
	Status int `form:"status" json:"status"`
}

func (s *UpdateTask) Run() *serialization.Result {
	var task model.Task
	count := 0
	model.Database.Model(model.Task{}).Where("id=? and user_id=?", s.TaskID, s.UserID).Count(&count).Find(&task)
	if count < 1 {
		return &serialization.Result{
			Code: 1,
			Message: "task not found",
		}
	}

	task.Title = s.Title
	task.Content = s.Content
	task.Status = s.Status
	err := model.Database.Save(&task).Error
	if err != nil {
		return &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("failed to update task: %s", err.Error()),
		}
	}

	return &serialization.Result{
		Code: 0,
		Message: "ok",
		Data: serialization.SerializeTask(task),
	}
}

type DeleteTask struct {
	UserID int
	TaskID int `form:"taskId" json:"taskId" binding:"required"`
}

func (s *DeleteTask) Run() *serialization.Result {
	var task model.Task
	count := 0
	model.Database.Model(model.Task{}).Where("id=? and user_id=?", s.TaskID, s.UserID).Count(&count).Find(&task)
	if count < 1 {
		return &serialization.Result{
			Code: 1,
			Message: "task not found",
		}
	}

	err := model.Database.Delete(&task).Error
	if err != nil {
		return &serialization.Result{
			Code: 1,
			Message: fmt.Sprintf("failed to update task: %s", err.Error()),
		}
	}

	return &serialization.Result{}
}

type SearchTask struct {
	UserID int
	Keyword string
	PageNo int
	PageSize int
}

func (s *SearchTask) Run() *serialization.Result {
	var tasks []model.Task
	total := 0
	if s.Keyword == "" {
		model.Database.
			Model(model.Task{}).
			Where("user_id=?", s.UserID).
			Count(&total).
			Limit(s.PageSize).Offset((s.PageNo-1)*s.PageSize).
			Find(&tasks)
	} else {
		model.Database.
			Model(model.Task{}).
			Where("user_id=?", s.UserID).
			Where("title like ? or content like ?", "%"+s.Keyword+"%", "%"+s.Keyword+"%").
			Count(&total).
			Limit(s.PageSize).Offset((s.PageNo-1)*s.PageSize).
			Find(&tasks)
	}
	
	return &serialization.Result{
		Data: serialization.SerializeTasks(tasks),
		Page: s.PageNo,
		Total: total,
	}
}
