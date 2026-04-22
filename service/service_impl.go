package service

import (
	"CRUD_API_PROJ/models"
	"CRUD_API_PROJ/repository"
	"errors"

	"gorm.io/gorm"
)

const (
	emptyTaskIDError  = "task ID cannot be empty"
	taskNotFoundError = "task not found"
)

type serviceConcrete struct {
	DB gorm.DB
}

func NewServiceConcrete() TaskService {
	return &serviceConcrete{
		DB: *repository.GetDB(),
	}
}

func (s *serviceConcrete) GetTask(taskID string) (*models.Task, error) {
	if taskID == "" {
		return nil, errors.New(emptyTaskIDError)
	}
	task, exists := repository.GetTaskByID(taskID)
	if !exists {
		return nil, errors.New(taskNotFoundError)
	}
	return &models.Task{
		TaskID:          task.TaskID,
		TaskName:        task.TaskName,
		TaskDescription: task.TaskDescription,
		TaskStatus:      task.TaskStatus,
		UserName:        task.UserName,
	}, nil
}

func (s *serviceConcrete) GetAllTasks() ([]*models.Task, error) {

	tasks, err := repository.GetAllTasks()
	if err != nil {
		return nil, err
	}

	result := make([]*models.Task, 0, len(tasks))

	for _, task := range tasks {
		result = append(result, &models.Task{
			TaskID:          task.TaskID,
			TaskName:        task.TaskName,
			TaskDescription: task.TaskDescription,
			TaskStatus:      task.TaskStatus,
			UserName:        task.UserName,
		})
	}

	return result, nil
}

func (s *serviceConcrete) CreateTask(request models.Task) (*models.Task, error) {
	if request.TaskID == "" {
		return nil, errors.New(emptyTaskIDError)
	}
	repository.CreateTask(&repository.Task{
		TaskID:          request.TaskID,
		TaskName:        request.TaskName,
		TaskDescription: request.TaskDescription,
		TaskStatus:      request.TaskStatus,
		UserName:        request.UserName,
	})
	return &request, nil
}

func (s *serviceConcrete) UpdateTask(request models.Task) (*models.Task, error) {
	if request.TaskID == "" {
		return nil, errors.New(emptyTaskIDError)
	}
	task, exists := repository.GetTaskByID(request.TaskID)
	if !exists {
		return nil, errors.New(taskNotFoundError)
	}
	task.TaskName = request.TaskName
	task.TaskDescription = request.TaskDescription
	task.TaskStatus = request.TaskStatus
	task.UserName = request.UserName
	repository.UpdateTask(task)
	return &request, nil
}

func (s *serviceConcrete) PatchTask(taskID string, request models.Task) (*models.Task, error) {
	if taskID == "" {
		return nil, errors.New(emptyTaskIDError)
	}
	task, exists := repository.GetTaskByID(taskID)
	if !exists {
		return nil, errors.New(taskNotFoundError)
	}
	if request.TaskName != "" {
		task.TaskName = request.TaskName
	}
	if request.TaskDescription != "" {
		task.TaskDescription = request.TaskDescription
	}
	if request.TaskStatus != "" {
		task.TaskStatus = request.TaskStatus
	}
	if request.UserName != "" {
		task.UserName = request.UserName
	}
	repository.UpdateTask(task)
	return &models.Task{
		TaskID:          task.TaskID,
		TaskName:        task.TaskName,
		TaskDescription: task.TaskDescription,
		TaskStatus:      task.TaskStatus,
		UserName:        task.UserName,
	}, nil
}

func (s *serviceConcrete) DeleteTask(taskID string) error {
	if taskID == "" {
		return errors.New(emptyTaskIDError)
	}
	task, exists := repository.GetTaskByID(taskID)
	if !exists {
		return errors.New(taskNotFoundError)
	}
	repository.DeleteTask(task.TaskID)
	return nil
}
