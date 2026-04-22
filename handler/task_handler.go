package handler

import (
	"CRUD_API_PROJ/models"
	"CRUD_API_PROJ/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	taskService service.TaskService
}

func NewHandler(taskService service.TaskService) *Handler {
	return &Handler{
		taskService: taskService,
	}
}

func (h *Handler) GetTask(c *gin.Context) {
	taskID := c.Param("id")
	task, err := h.taskService.GetTask(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *Handler) GetAllTasks(c *gin.Context) {
	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var taskRequest models.Task
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := h.taskService.CreateTask(taskRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	var taskRequest models.Task
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask, err := h.taskService.UpdateTask(taskRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

func (h *Handler) PatchTask(c *gin.Context) {
	taskID := c.Param("id")
	var taskRequest models.Task
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask, err := h.taskService.PatchTask(taskID, taskRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}
func (h *Handler) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	err := h.taskService.DeleteTask(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
