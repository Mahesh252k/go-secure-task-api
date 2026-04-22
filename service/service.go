package service

import (
	"CRUD_API_PROJ/models"
)

type TaskService interface {
	GetTask(taskID string) (*models.Task, error)
	GetAllTasks() ([]*models.Task, error)
	CreateTask(request models.Task) (*models.Task, error)
	UpdateTask(request models.Task) (*models.Task, error)
	PatchTask(taskID string, request models.Task) (*models.Task, error)
	DeleteTask(taskID string) error
}
