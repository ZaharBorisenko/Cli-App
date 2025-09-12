package handlers

import (
	"fmt"
	"github.com/ZaharBorisenko/Cli-App/models"
)

type PriorityManager struct {
	todos *Todos
}

func NewPriorityManager(todos *Todos) *PriorityManager {
	return &PriorityManager{todos: todos}
}

func (pm *PriorityManager) SetPriority(id int, priority models.Priority) error {
	if err := (*pm.todos).ValidateId(id); err != nil {
		return err
	}

	(*pm.todos)[id].Priority = priority
	return nil
}

func (pm *PriorityManager) ValidatePriority(priorityStr string) (models.Priority, error) {
	switch priorityStr {
	case "high", "h":
		return models.PriorityHigh, nil
	case "medium", "med", "m":
		return models.PriorityMedium, nil
	case "low", "l":
		return models.PriorityLow, nil
	case "none", "":
		return models.PriorityNone, nil
	default:
		return models.PriorityNone, fmt.Errorf("invalid priority: %s. Use: high, medium, low, none", priorityStr)
	}
}

func (pm *PriorityManager) PrintByPriority(priority models.Priority) {
	var filteredTodos []models.Todo
	found := false

	for _, t := range *pm.todos {
		if t.Priority == priority {
			filteredTodos = append(filteredTodos, t)
			found = true
		}
	}

	if !found {
		fmt.Printf("No tasks found with priority '%s'\n", priority)
		return
	}

	PrintTable(filteredTodos, TableConfig{
		ShowCategory:    true,
		ShowCompletedAt: false,
		ShowPriority:    true,
		ShowStatus:      true,
	})
}
