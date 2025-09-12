package handlers

import (
	"fmt"
	"github.com/ZaharBorisenko/Cli-App/models"
	table2 "github.com/aquasecurity/table"
	"os"
	"strconv"
	"time"
)

type TableConfig struct {
	ShowCategory    bool
	ShowCompletedAt bool
	ShowPriority    bool
	ShowStatus      bool
}

func PrintTable(todos []models.Todo, config TableConfig) {
	table := table2.New(os.Stdout)
	table.SetRowLines(false)

	statusManager := StatusManager{} // Для получения символов статусов

	// Динамически формируем заголовки
	headers := []string{"#", "Title", "Description"}
	if config.ShowCategory {
		headers = append(headers, "Category")
	}
	if config.ShowPriority {
		headers = append(headers, "Priority")
	}
	if config.ShowStatus {
		headers = append(headers, "Status")
	}
	headers = append(headers, "Created at")
	if config.ShowCompletedAt {
		headers = append(headers, "Completed at")
	}

	table.SetHeaders(headers...)

	for id, t := range todos {
		// Формируем строку
		row := []string{
			strconv.Itoa(id),
			t.Title,
			t.Description,
		}

		if config.ShowCategory {
			row = append(row, t.Category)
		}

		if config.ShowPriority {
			var priorityDisplay string
			switch t.Priority {
			case models.PriorityHigh:
				priorityDisplay = "🔴 high"
			case models.PriorityMedium:
				priorityDisplay = "🟡 medium"
			case models.PriorityLow:
				priorityDisplay = "🟢 low"
			default:
				priorityDisplay = "⚪ none"
			}
			row = append(row, priorityDisplay)
		}

		if config.ShowStatus {
			statusSymbol := statusManager.GetStatusSymbol(t.Status)
			row = append(row, fmt.Sprintf("%s %s", statusSymbol, t.Status))
		}

		row = append(row, t.CreatedAt.Format(time.RFC1123))

		if config.ShowCompletedAt && t.CompletedAt != nil {
			row = append(row, t.CompletedAt.Format(time.RFC1123))
		} else if config.ShowCompletedAt {
			row = append(row, "") // Пустая строка если нет времени завершения
		}

		table.AddRow(row...)
	}

	table.Render()
}
