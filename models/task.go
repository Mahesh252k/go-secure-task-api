package models

type Task struct {
	TaskID          string `json:"taskId"`
	TaskName        string `json:"taskName"`
	TaskDescription string `json:"taskDescription"`
	TaskStatus      string `json:"taskStatus"`
	UserName        string `json:"userName"`
}

func NewTask(taskID, taskName, taskDescription, taskStatus, userName string) *Task {
	return &Task{
		TaskID:          taskID,
		TaskName:        taskName,
		TaskDescription: taskDescription,
		TaskStatus:      taskStatus,
		UserName:        userName,
	}
}
