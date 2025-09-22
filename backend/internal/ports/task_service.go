package ports

import "github.com/PrinceM13/daily-helper/internal/domain"

type TaskService interface {
	GetTasks(page, limit int) ([]domain.Task, int, error)
	GetTaskByID(id int) (*domain.Task, error)
	CreateTask(task *domain.Task) (*domain.Task, error)
	UpdateTask(task *domain.Task) (*domain.Task, error)
	DeleteTask(id int) (*domain.Task, error)
}
