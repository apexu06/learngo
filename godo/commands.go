package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func (a *App) TaskAddCommand(args []string) {
	var newTask Task

	if len(args) <= 1 {
		log.Fatalln("KEIN ARGS DU ARSCHLOCH")
	}

	newTask.Id = len(a.Tasks) + 1
	newTask.Content = args[1]
	newTask.State = "todo"
	newTask.DateCreated = time.Now()

	a.Tasks = append(a.Tasks, newTask)

}

func TaskDeleteCommand(args []string) {

}

func TaskUpdateCommand(args []string) {

}

func (a *App) TaskListCommand(args []string) {
	longest := len("Description")
	for _, t := range a.Tasks {
		longest = max(longest, len(t.Content))
	}

	fmt.Printf("ID | %-*s | Status\n", longest, "Description")

	fmt.Printf("---+-%-*s-+---------\n", longest, strings.Repeat("-", longest))

	for _, task := range a.Tasks {
		fmt.Printf("%2d | %-*s | %s\n", task.Id, longest, task.Content, task.State)
	}
}
