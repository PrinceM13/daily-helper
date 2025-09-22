package application

import (
	"github.com/PrinceM13/daily-helper/internal/domain"
	"github.com/PrinceM13/daily-helper/internal/ports"
)

type TaskService struct {
	repo ports.TaskRepository
}

func NewTaskService(r ports.TaskRepository) ports.TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) GetTasks(page, limit int) ([]domain.Task, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	tasks, total, err := s.repo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return tasks, total, nil
}

func (s *TaskService) GetTaskByID(id int) (*domain.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *domain.Task) (*domain.Task, error) {
	// Add business rules here if needed before updating
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id int) (*domain.Task, error) {
	// Add business rules here if needed before deleting
	return s.repo.Delete(id)
}

func (s *TaskService) CreateTask(task *domain.Task) (*domain.Task, error) {
	// Add validation or other rules here if needed before creating
	return s.repo.Create(task)
}
