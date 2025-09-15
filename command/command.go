package command

import (
	"flag"
	"fmt"
	"github.com/ZaharBorisenko/Cli-App/handlers"
	"github.com/ZaharBorisenko/Cli-App/models"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add              string
	AddDesc          string
	Del              int
	Edit             string
	Toggle           int
	List             bool
	AddCategory      string
	SetCategory      string
	ListByCat        string
	ListCats         bool
	SetPriority      string
	ListByPriority   string
	Colors           bool
	SetStatus        string
	ListByStatus     string
	ListDone         bool
	ListActive       bool
	StatisticTodo    bool
	SetTimeCompleted string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	//task
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.StringVar(&cf.AddDesc, "addDesc", "", "Add description a todo by index & specify a description. id:new_desc")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.BoolVar(&cf.StatisticTodo, "stats", false, "Stats all todos")

	// category
	flag.StringVar(&cf.AddCategory, "addCategory", "", "Add a new category")
	flag.StringVar(&cf.SetCategory, "setCategory", "", "Set category for a task. id:category")
	flag.StringVar(&cf.ListByCat, "listByCat", "", "List tasks by category")
	flag.BoolVar(&cf.ListCats, "listCats", false, "List all categories")

	//priority
	flag.StringVar(&cf.SetPriority, "setPriority", "", "Set priority for a task. id:priority (high/medium/low/none)")
	flag.StringVar(&cf.ListByPriority, "listByPriority", "", "List tasks by priority (high/medium/low/none)")
	flag.BoolVar(&cf.Colors, "colors", true, "Enable/disable colors")

	//status
	flag.StringVar(&cf.SetStatus, "setStatus", "", "Set status for a task. id:status (todo/inprogress/done)")
	flag.StringVar(&cf.ListByStatus, "listByStatus", "", "List tasks by status (todo/inprogress/done)")
	flag.BoolVar(&cf.ListDone, "listDone", false, "List completed tasks")
	flag.BoolVar(&cf.ListActive, "listActive", false, "List active tasks (not completed)")

	//time
	flag.StringVar(&cf.SetTimeCompleted, "setTime", "", "set date completed")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *handlers.Todos) {
	categoryManager := handlers.NewCategoryManager(todos)
	priorityManager := handlers.NewPriorityManager(todos)
	statusManager := handlers.NewStatusManager(todos)
	timeManager := handlers.NewTimeManager(todos)

	switch {
	case cf.SetTimeCompleted != "":
		parts := strings.SplitN(cf.SetTimeCompleted, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid format for setTimeCompleted")
			return
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index")
			return
		}
		if err := timeManager.SetDateCompleted(index, parts[1]); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("time '%s' set for task %d\n", parts[1], index)
			todos.PrintTodos()
		}
	case cf.List:
		todos.PrintTodos()
	case cf.StatisticTodo:
		todos.StatisticTodo()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid format for edit")
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index")
		}
		todos.Edit(index, parts[1])
	case cf.AddDesc != "":
		parts := strings.SplitN(cf.AddDesc, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid format for desc")
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index")
		}
		todos.AddDescription(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Del != -1:
		todos.Delete(cf.Del)
		//category
	case cf.AddCategory != "":
		fmt.Printf("Category '%s' is ready to use\n", cf.AddCategory)

	case cf.SetCategory != "":
		parts := strings.SplitN(cf.SetCategory, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid format for setCategory. Use: id:category")
			return
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index")
			return
		}
		if err := categoryManager.SetCategory(index, parts[1]); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Category '%s' set for task %d\n", parts[1], index)
			todos.PrintTodos()
		}

	case cf.ListByCat != "":
		categoryManager.PrintByCategory(cf.ListByCat)

	case cf.ListCats:
		categoryManager.PrintAllCategories()
	case cf.SetPriority != "":
		parts := strings.SplitN(cf.SetPriority, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid format for setPriority. Use: id:priority")
			return
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index")
			return
		}
		priority, err := priorityManager.ValidatePriority(strings.ToLower(parts[1]))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if err := priorityManager.SetPriority(index, priority); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Priority '%s' set for task %d\n", priority, index)
			todos.PrintTodos()
		}

	case cf.ListByPriority != "":
		priority, err := priorityManager.ValidatePriority(strings.ToLower(cf.ListByPriority))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		priorityManager.PrintByPriority(priority)
	case cf.SetStatus != "":
		parts := strings.SplitN(cf.SetStatus, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid format for setStatus. Use: id:status")
			return
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index")
			return
		}
		status, err := statusManager.ValidateStatus(strings.ToLower(parts[1]))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if err := statusManager.SetStatus(index, status); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Status '%s' set for task %d\n", status, index)
			todos.PrintTodos()
		}

	case cf.ListByStatus != "":
		status, err := statusManager.ValidateStatus(strings.ToLower(cf.ListByStatus))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		statusManager.PrintByStatus(status)

	case cf.ListDone:
		statusManager.PrintByStatus(models.StatusDone)

	case cf.ListActive:
		var activeTodos []models.Todo
		for _, t := range *todos {
			if t.Status != models.StatusDone {
				activeTodos = append(activeTodos, t)
			}
		}
		handlers.PrintTable(activeTodos, handlers.TableConfig{
			ShowCategory: true,
			ShowPriority: true,
			ShowStatus:   true,
		})

	default:
		fmt.Println("invalid cmd command")
	}
}
