package main

import (
	"fmt"
	"log"
	"os"
)

func TaskAddCommand(args []string) {

}

func TaskDeleteCommand(args []string) {

}

func TaskUpdateCommand(args []string) {

}

func (a *App) TaskListCommand(args []string) {
	if len(os.Args) > 2 {
		log.Fatalln("Unrecognized argument: ", args[1])
	}

	for _, task := range a.Tasks {
		fmt.Printf("%d - %s - %s\n", task.Id, task.Content, task.State)
	}

}
