package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	TaskID          string `gorm:"primaryKey"`
	TaskName        string `gorm:"type:varchar(255)"`
	TaskDescription string `gorm:"type:text"`
	TaskStatus      string `gorm:"type:varchar(50)"`
	UserName        string `gorm:"type:varchar(255)"`
}

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	fmt.Println("Database connection established")
}

func CreateTable() {
	err := DB.AutoMigrate(&Task{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully")
}

func GetDB() *gorm.DB {
	return DB
}

func CreateTask(task *Task) error {
	result := DB.Create(task)
	return result.Error
}

func GetTaskByID(taskID string) (*Task, bool) {
	var task Task
	result := DB.First(&task, "task_id = ?", taskID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, false
		}
		return nil, false
	}
	return &task, true
}

func GetAllTasks() ([]*Task, error) {
	var tasks []*Task
	result := DB.Find(&tasks)
	return tasks, result.Error
}

func UpdateTask(task *Task) error {
	result := DB.Save(task)
	return result.Error
}

func DeleteTask(taskID string) error {
	result := DB.Delete(&Task{}, "task_id = ?", taskID)
	return result.Error
}
