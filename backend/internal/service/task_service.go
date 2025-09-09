package service

import "github.com/PrinceM13/daily-helper/internal/domain"

type TaskService struct {
	Repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) GetTasks(page, limit int) ([]domain.Task, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	tasks, total, err := s.Repo.GetPaginated(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return tasks, total, nil
}

func (s *TaskService) GetTaskByID(id int) (*domain.Task, error) {
	return s.Repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *domain.Task) error {
	// Add business rules here if needed before updating
	return s.Repo.Update(task)
}

func (s *TaskService) DeleteTask(id int) error {
	// Add business rules here if needed before deleting
	return s.Repo.Delete(id)
}

func (s *TaskService) CreateTask(task *domain.Task) error {
	// Add validation or other rules here if needed before creating
	return s.Repo.Create(task)
}
