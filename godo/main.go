package main

import (
	"fmt"
	"log"
	"os"
)

func printHelp() {
	// TODO
}

type App struct {
	Tasks []Task
	Path  string
}

func InitApp(path string) (*App, error) {
	tasks, err := GetAllTasks(path)
	if err != nil {
		return nil, err
	}

	return &App{
		Tasks: tasks,
		Path:  path,
	}, nil
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	app, err := InitApp("tasks.json")

	if err != nil {
		log.Fatal(err)
	}

	app.AddTask()
	if err := run(app, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(app *App, args []string) error {
	switch operation := args[0]; operation {
	case "add":
	case "delete":
	case "update":
	case "list":
		app.TaskListCommand(args)
	default:
		return fmt.Errorf("Unknown command: %s", args[0])
	}
	return nil
}
