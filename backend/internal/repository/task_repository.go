package repository

import (
	"errors"

	"github.com/PrinceM13/daily-helper/internal/domain"
)

type TaskRepo struct {
	tasks []domain.Task
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		// tasks: []domain.Task{}, // Initialize with an empty slice
		tasks: []domain.Task{
			{ID: 1, Title: "Alice Buy groceries", Completed: false},
			{ID: 2, Title: "Read a book", Completed: true},
			{ID: 3, Title: "Write Go code", Completed: false},
			{ID: 4, Title: "Exercise", Completed: true},
			{ID: 5, Title: "Call mom", Completed: false},
			{ID: 6, Title: "Pay bills", Completed: true},
			{ID: 7, Title: "Clean the house", Completed: false},
			{ID: 8, Title: "Walk the dog", Completed: true},
			{ID: 9, Title: "Cook dinner", Completed: false},
			{ID: 10, Title: "Meditate", Completed: true},
			{ID: 11, Title: "Plan vacation", Completed: false},
			{ID: 12, Title: "Attend meeting", Completed: true},
			{ID: 13, Title: "Write blog post", Completed: false},
			{ID: 14, Title: "Fix bugs", Completed: true},
			{ID: 15, Title: "Update resume", Completed: false},
		},
	}
}

func (r *TaskRepo) GetAll() ([]domain.Task, error) {
	return r.tasks, nil
}

func (r *TaskRepo) GetByID(id int) (*domain.Task, error) {
	for _, t := range r.tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *TaskRepo) Create(task *domain.Task) error {
	r.tasks = append(r.tasks, *task)
	return nil
}

func (r *TaskRepo) Update(task *domain.Task) error {
	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = *task
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *TaskRepo) Delete(id int) error {
	for i, t := range r.tasks {
		if t.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *TaskRepo) GetPaginated(limit, offset int) ([]domain.Task, int, error) {
	total := len(r.tasks)

	// Guard offset
	if offset > total {
		return []domain.Task{}, total, nil
	}

	// End index (min of offset+limit and total)
	end := offset + limit
	if end > total {
		end = total
	}

	return r.tasks[offset:end], total, nil
}
