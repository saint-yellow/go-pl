package serialization

import (
	"time"

	"github.com/saint-yellow/go-pl/toys/memorandum/model"
)

type Task struct {
	ID uint `form:"id" json:"id"`
	CreatedAt time.Time `form:"cretedAt" json:"createdAt"`
	UpdatedAt time.Time `form:"updatedAt" json:"updatedAt"`
	UserID int `form:"userId" json:"userId"`
	Title string `form:"title" json:"title"`
	Status int `form:"status" json:"status"`
	Content string `form:"content" json:"content"`
	StartTime int64 `form:"startTime" json:"startTime"`
	EndTime int64 `form:"endTime" json:"endTime"`
	BrowseCount uint64 `form:"browseCount" json:"browseCount"`
}

func newTask(task model.Task) Task {
	return Task{
		ID: task.ID,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		UserID: task.UserID,
		Title: task.Title,
		Content: task.Content,
		Status: task.Status,
		StartTime: task.StartTime,
		EndTime: task.EndTime,
		BrowseCount: task.BrowseCount(),
	}
}

func SerializeTask(task model.Task) Task {
	return newTask(task)
}

func SerializeTasks(tasks []model.Task) []Task {
	result := make([]Task, len(tasks))
	for i := range tasks {
		result[i] = newTask(tasks[i])
	}
	return result
}