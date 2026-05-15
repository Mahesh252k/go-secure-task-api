package main

import (
	"CRUD_API_PROJ/handler"
	loginhandler "CRUD_API_PROJ/loginHandler"
	"CRUD_API_PROJ/middleware"
	"CRUD_API_PROJ/repository"
	"CRUD_API_PROJ/service"
	"fmt"

	serviceauth "CRUD_API_PROJ/serviceAuth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	route := gin.Default()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	repository.ConnectDB()
	repository.CreateUserTable()
	repository.CreateTable()
	taskService := service.NewServiceConcrete()
	taskHandler := handler.NewHandler(taskService)

	taskServiceAuth := serviceauth.NewServiceConcrete()
	loginHandler := loginhandler.NewLoginHandler(taskServiceAuth)
	route.POST("/register", loginHandler.Register)
	route.POST("/login", loginHandler.Login)

	protected := route.Group("/")

	protected.GET("/tasks", taskHandler.GetAllTasks)
	protected.GET("/task/:id", taskHandler.GetTask)
	protected.PUT("/task/:id", middleware.WritterMiddleware(), taskHandler.UpdateTask)
	protected.PATCH("/task/:id", middleware.WritterMiddleware(), taskHandler.PatchTask)
	protected.POST("/task", middleware.AdminMiddleware(), taskHandler.CreateTask)
	protected.DELETE("/task/:id", middleware.AdminMiddleware(), taskHandler.DeleteTask)
	route.Run(":8080")
}
