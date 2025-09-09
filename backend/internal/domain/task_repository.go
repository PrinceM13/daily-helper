package domain

type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id int) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id int) error
	GetPaginated(limit, offset int) ([]Task, int, error)
}
