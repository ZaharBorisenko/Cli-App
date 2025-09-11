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
	table := table2.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Description", "Completed", "Created at", "Completed at")
	for id, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(id), t.Title, t.Description, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}

func (todos *Todos) Add(title string) {
	newTodo := models.Todo{
		Title:       title,
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
