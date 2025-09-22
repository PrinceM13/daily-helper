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

func (s *TaskService) ListTasks(page, limit int) ([]domain.Task, int, error) {
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

func (s *TaskService) RetrieveTaskByID(id int) (*domain.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) AddNewTask(task *domain.Task) (*domain.Task, error) {
	// Add validation or other rules here if needed before creating
	return s.repo.Create(task)
}

func (s *TaskService) ModifyTask(task *domain.Task) (*domain.Task, error) {
	// Add business rules here if needed before updating
	return s.repo.Update(task)
}

func (s *TaskService) RemoveTask(id int) (*domain.Task, error) {
	// Add business rules here if needed before deleting
	return s.repo.Delete(id)
}
