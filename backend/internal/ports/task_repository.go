package ports

import "github.com/PrinceM13/daily-helper/internal/domain"

type TaskRepository interface {
	GetAll(int, int) ([]domain.Task, int, error)
	GetByID(id int) (*domain.Task, error)
	Create(task *domain.Task) (*domain.Task, error)
	Update(task *domain.Task) (*domain.Task, error)
	Delete(id int) (*domain.Task, error)
}
