package repository

import (
	"errors"

	"github.com/PrinceM13/daily-helper/internal/domain"
	"github.com/PrinceM13/daily-helper/internal/ports"
)

type InMemoryTaskRepo struct {
	tasks map[int]domain.Task
}

func strPtr(s string) *string { return &s }
func boolPtr(b bool) *bool    { return &b }

func NewInMemoryTaskRepo() ports.TaskRepository {
	repo := &InMemoryTaskRepo{tasks: make(map[int]domain.Task)}

	// Seed with some default data
	defaultTasks := []domain.Task{
		{ID: 1, Title: strPtr("Buy groceries"), Completed: boolPtr(false)},
		{ID: 2, Title: strPtr("Read a book"), Completed: boolPtr(true)},
		{ID: 3, Title: strPtr("Write Go code"), Completed: boolPtr(false)},
		{ID: 4, Title: strPtr("Exercise"), Completed: boolPtr(true)},
		{ID: 5, Title: strPtr("Call mom"), Completed: boolPtr(false)},
	}

	for _, t := range defaultTasks {
		repo.tasks[t.ID] = t
	}

	return repo
}

func (r *InMemoryTaskRepo) GetAll(page, limit int) ([]domain.Task, int, error) {
	list := []domain.Task{}
	for _, t := range r.tasks {
		list = append(list, t)
	}

	total := len(list)

	start := (page - 1) * limit
	if start > total {
		return []domain.Task{}, total, nil
	}

	end := start + limit
	if end > total {
		end = total
	}

	return list[start:end], total, nil
}

func (r *InMemoryTaskRepo) GetByID(id int) (*domain.Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

func (r *InMemoryTaskRepo) Create(task *domain.Task) (*domain.Task, error) {
	maxID := 0

	for id := range r.tasks {
		if id > maxID {
			maxID = id
		}
	}
	task.ID = maxID + 1

	r.tasks[task.ID] = *task

	return task, nil
}

func (r *InMemoryTaskRepo) Update(task *domain.Task) (*domain.Task, error) {
	existing, ok := r.tasks[task.ID]
	if !ok {
		return nil, errors.New("task not found")
	}

	if task.Title != nil {
		existing.Title = task.Title
	}
	if task.Completed != nil {
		existing.Completed = task.Completed
	}

	r.tasks[task.ID] = existing
	return &existing, nil
}

func (r *InMemoryTaskRepo) Delete(id int) (*domain.Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}

	delete(r.tasks, id)

	return &task, nil
}
