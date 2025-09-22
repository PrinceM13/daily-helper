// File: /internal/repository/factory.go

package repository

import (
	"fmt"
	"os"

	"github.com/PrinceM13/daily-helper/internal/adapters/repository/inmemory"
	"github.com/PrinceM13/daily-helper/internal/adapters/repository/postgres"
	"github.com/PrinceM13/daily-helper/internal/ports"
)

func NewTaskRepo() (ports.TaskRepository, error) {
	repoType := os.Getenv("TASK_REPO") // "inmemory" or "postgres" for now

	// Default to "inmemory"
	if repoType == "" {
		repoType = "inmemory"
	}

	switch repoType {
	case "inmemory":
		return inmemory.NewInMemoryTaskRepo(), nil
	case "postgres":
		// In a real app, you would connect to the database here.
		db := &postgres.PostgresDB{}
		return postgres.NewPostgresTaskRepo(db), nil
	default:
		return nil, fmt.Errorf("unknown repo type: %s", repoType)
	}
}
