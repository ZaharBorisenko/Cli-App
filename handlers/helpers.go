package handlers

import (
	"github.com/ZaharBorisenko/Cli-App/models"
	table2 "github.com/aquasecurity/table"
	"os"
	"strconv"
	"time"
)

type TableConfig struct {
	ShowCategory    bool
	ShowCompletedAt bool
}

func PrintTable(todos []models.Todo, config TableConfig) {
	table := table2.New(os.Stdout)
	table.SetRowLines(false)

	// Динамически формируем заголовки
	headers := []string{"#", "Title", "Description"}
	if config.ShowCategory {
		headers = append(headers, "Category")
	}
	headers = append(headers, "Completed", "Created at")
	if config.ShowCompletedAt {
		headers = append(headers, "Completed at")
	}

	table.SetHeaders(headers...)

	for id, t := range todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		row := []string{
			strconv.Itoa(id),
			t.Title,
			t.Description,
		}

		if config.ShowCategory {
			row = append(row, t.Category)
		}

		row = append(row, completed, t.CreatedAt.Format(time.RFC1123))

		if config.ShowCompletedAt {
			row = append(row, completedAt)
		}

		table.AddRow(row...)
	}

	table.Render()
}
