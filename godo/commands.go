package main

import (
	"fmt"
	"strings"
)

func TaskAddCommand(args []string) {

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
