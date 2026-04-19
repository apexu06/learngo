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

	longest := 0
	for _, t := range a.Tasks {
		if len(t.Content) > longest {
			longest = len(t.Content)
		}
	}

	longest = max(len("Description"), longest)
	header := fmt.Sprintf("ID | %-*s | Status\n", longest, "Description")

	for range len(header) {
		fmt.Printf("-")
	}

	for _, task := range a.Tasks {
		fmt.Printf("%2d | %-*s | %s\n", task.Id, longest, task.Content, task.State)
	}

}
