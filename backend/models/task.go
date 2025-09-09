package models

import (
	"errors"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []Task{
	{ID: 1, Title: "Buy groceries", Completed: false},
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
}

func GetAllTasks() ([]Task, error) {
	return tasks, nil
}

func GetTaskByID(id int) (*Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (t *Task) Update() error {
	for i, task := range tasks {
		if task.ID == t.ID {
			tasks[i] = *t
			return nil
		}
	}
	return errors.New("task not found")
}

func (t *Task) Delete() error {
	for i, task := range tasks {
		if task.ID == t.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
