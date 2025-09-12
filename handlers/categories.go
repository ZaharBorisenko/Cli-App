package handlers

import (
	"fmt"
	"github.com/ZaharBorisenko/Cli-App/models"
	table2 "github.com/aquasecurity/table"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CategoryManager struct {
	todos *Todos
}

func NewCategoryManager(todos *Todos) *CategoryManager {
	return &CategoryManager{todos: todos}
}

func (cm *CategoryManager) GetUniqueCategories() []string {
	categories := make(map[string]bool)
	for _, todo := range *cm.todos {
		if todo.Category != "" {
			categories[todo.Category] = true
		}
	}

	result := make([]string, 0, len(categories))
	for category := range categories {
		result = append(result, category)
	}
	sort.Strings(result)
	return result
}

func (cm *CategoryManager) SetCategory(id int, category string) error {
	if err := (*cm.todos).ValidateId(id); err != nil {
		return err
	}

	(*cm.todos)[id].Category = strings.TrimSpace(category)
	return nil
}

func (cm *CategoryManager) PrintByCategory(category string) {
	var filteredTodos []models.Todo
	found := false

	for _, t := range *cm.todos {
		if t.Category == category {
			filteredTodos = append(filteredTodos, t)
			found = true
		}
	}

	if !found {
		fmt.Printf("No tasks found in category '%s'\n", category)
		return
	}

	PrintTable(filteredTodos, TableConfig{
		ShowCategory:    true,
		ShowCompletedAt: false, // Можно изменить на true если нужно
	})
}

func (cm *CategoryManager) PrintAllCategories() {
	categories := cm.GetUniqueCategories()
	if len(categories) == 0 {
		fmt.Println("No categories found")
		return
	}

	table := table2.New(os.Stdout)
	table.SetHeaders("Category", "Task Count")

	for _, category := range categories {
		count := cm.CountTasksInCategory(category)
		table.AddRow(category, strconv.Itoa(count))
	}

	table.Render()
}

func (cm *CategoryManager) CountTasksInCategory(category string) int {
	count := 0
	for _, todo := range *cm.todos {
		if todo.Category == category {
			count++
		}
	}
	return count
}
