package main

import (
	"github.com/ZaharBorisenko/Cli-App/command"
	"github.com/ZaharBorisenko/Cli-App/handlers"
	storage2 "github.com/ZaharBorisenko/Cli-App/storage"
)

func main() {
	todos := handlers.Todos{}
	storage := storage2.NewStorage[handlers.Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := command.NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
