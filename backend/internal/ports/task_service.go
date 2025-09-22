package ports

import "github.com/PrinceM13/daily-helper/internal/domain"

type TaskService interface {
	ListTasks(page, limit int) ([]domain.Task, int, error)
	RetrieveTaskByID(id int) (*domain.Task, error)
	AddNewTask(task *domain.Task) (*domain.Task, error)
	ModifyTask(task *domain.Task) (*domain.Task, error)
	RemoveTask(id int) (*domain.Task, error)
}
