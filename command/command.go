package command

import (
	"flag"
	"fmt"
	"github.com/ZaharBorisenko/Cli-App/handlers"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add         string
	AddDesc     string
	Del         int
	Edit        string
	Toggle      int
	List        bool
	AddCategory string
	SetCategory string
	ListByCat   string
	ListCats    bool
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

	// category
	flag.StringVar(&cf.AddCategory, "addCategory", "", "Add a new category")
	flag.StringVar(&cf.SetCategory, "setCategory", "", "Set category for a task. id:category")
	flag.StringVar(&cf.ListByCat, "listByCat", "", "List tasks by category")
	flag.BoolVar(&cf.ListCats, "listCats", false, "List all categories")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *handlers.Todos) {
	categoryManager := handlers.NewCategoryManager(todos)
	switch {
	case cf.List:
		todos.PrintTodos()
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

	default:
		fmt.Println("invalid cmd command")
	}
}

// go run main.go -list
// go run main.go -edit "2:ALLO SUKA"
// go run main.go -add "test"
// go run main.go -del 0
// go run main.go -addDesc "2:ALLO SUKAdddd"
