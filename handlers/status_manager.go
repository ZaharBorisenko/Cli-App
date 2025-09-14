package handlers

import (
	"fmt"
	"github.com/ZaharBorisenko/Cli-App/models"
	"time"
)

type StatusManager struct {
	todos *Todos
}

func NewStatusManager(todos *Todos) *StatusManager {
	return &StatusManager{todos: todos}
}

func (sm *StatusManager) SetStatus(id int, status models.Status) error {
	if err := (*sm.todos).ValidateId(id); err != nil {
		return err
	}

	todo := &(*sm.todos)[id]
	todo.Status = status
	todo.UpdatedAt = time.Now()

	if status == models.StatusDone && !todo.Completed {
		completedTime := time.Now()
		todo.CompletedAt = &completedTime
		todo.Completed = true
	} else if status != models.StatusDone && todo.Completed {
		todo.CompletedAt = nil
		todo.Completed = false
	}

	return nil
}

func (sm *StatusManager) ValidateStatus(statusStr string) (models.Status, error) {
	switch statusStr {
	case "todo", "t":
		return models.StatusTodo, nil
	case "inprogress", "ip", "progress":
		return models.StatusInProgress, nil
	case "done", "d":
		return models.StatusDone, nil
	default:
		return models.StatusTodo, fmt.Errorf("invalid status: %s. Use: todo, inprogress, done", statusStr)
	}
}

func (sm *StatusManager) GetStatusSymbol(status models.Status) string {
	switch status {
	case models.StatusTodo:
		return "⏳"
	case models.StatusInProgress:
		return "🚧"
	case models.StatusDone:
		return "✅"
	default:
		return "❓"
	}
}

func (sm *StatusManager) PrintByStatus(status models.Status) {
	var filteredTodos []models.Todo

	for _, t := range *sm.todos {
		if t.Status == status {
			filteredTodos = append(filteredTodos, t)
		}
	}

	if len(filteredTodos) == 0 {
		fmt.Printf("No tasks found with status '%s'\n", status)
		return
	}

	PrintTable(filteredTodos, TableConfig{
		ShowCategory:    true,
		ShowCompletedAt: status == models.StatusDone,
		ShowPriority:    true,
		ShowStatus:      true,
	})
}
