package handlers

import (
	"errors"
	"fmt"
	"github.com/ZaharBorisenko/Cli-App/models"
	table2 "github.com/aquasecurity/table"
	"os"
	"strconv"
	"time"
)

type Todos []models.Todo

func (todos *Todos) PrintTodos() {
	PrintTable(*todos, TableConfig{
		ShowCategory:    true,
		ShowCompletedAt: true,
		ShowPriority:    true,
		ShowStatus:      true,
	})
}

func (todos *Todos) Add(title string) {
	newTodo := models.Todo{
		Title:       title,
		Status:      models.StatusTodo,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, newTodo)
	todos.PrintTodos()
}

func (todos *Todos) AddDescription(id int, description string) error {
	t := *todos
	if err := t.ValidateId(id); err != nil {
		return err
	}
	t[id].Description = description
	todos.PrintTodos()
	return nil
}

func (todos *Todos) ValidateId(id int) error {
	err := errors.New("invalid id")

	if id < 0 || id >= len(*todos) {
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) Delete(id int) error {
	t := *todos

	if err := t.ValidateId(id); err != nil {
		return err
	}

	*todos = append(t[:id], t[id+1:]...)
	todos.PrintTodos()
	return nil
}

func (todos *Todos) Toggle(id int) error {
	t := *todos

	if err := t.ValidateId(id); err != nil {
		return err
	}

	isCompleted := t[id].Completed

	if !isCompleted {
		CompletedTime := time.Now()
		t[id].CompletedAt = &CompletedTime
	} else {
		t[id].CompletedAt = nil
	}

	t[id].Completed = !isCompleted
	*todos = t
	todos.PrintTodos()
	return nil
}

func (todos *Todos) Edit(id int, newTitle string) error {
	t := *todos

	if err := t.ValidateId(id); err != nil {
		return err
	}

	t[id].Title = newTitle
	*todos = t
	todos.PrintTodos()
	return nil
}

func (todos *Todos) StatisticTodo() {
	if len(*todos) == 0 {
		fmt.Println("No tasks available")
		return
	}

	var (
		allTask    = len(*todos)
		doneTask   = 0
		todoTask   = 0
		inProgress = 0
	)

	for _, task := range *todos {
		switch task.Status {
		case models.StatusDone:
			doneTask++
		case models.StatusTodo:
			todoTask++
		case models.StatusInProgress:
			inProgress++
		}
	}

	table := table2.New(os.Stdout)
	table.SetHeaders("Metric", "Count", "Percentage")

	table.AddRow("Total Tasks", strconv.Itoa(allTask), "100%")
	table.AddRow("✅ Done", strconv.Itoa(doneTask),
		fmt.Sprintf("%.1f%%", float64(doneTask)/float64(allTask)*100))
	table.AddRow("🚧 In Progress", strconv.Itoa(inProgress),
		fmt.Sprintf("%.1f%%", float64(inProgress)/float64(allTask)*100))
	table.AddRow("⏳ Todo", strconv.Itoa(todoTask),
		fmt.Sprintf("%.1f%%", float64(todoTask)/float64(allTask)*100))

	fmt.Println("📊 Task Statistics")
	fmt.Println("==================")
	table.Render()
}
