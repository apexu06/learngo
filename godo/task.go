package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	DateCreated time.Time `json:"date_created"`
	State       string    `json:"state"`
}

func GetAllTasks(path string) ([]Task, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to open file: %s", err)
	}

	var tasks []Task
	err = json.Unmarshal(content, &tasks)
	if err != nil {
		return nil, fmt.Errorf("Could not parse json: %s", err)
	}

	return tasks, nil
}

func (a *App) AddTask() {

	// for _, v := range a.Tasks {
	// 	println("TASKSSSS")
	// 	println(v.Content)
	// }

	jsonData, err := json.Marshal(a.Tasks)

	if err != nil {
		log.Fatalln("Error with marshal", err)
	}

	f, err := os.OpenFile("./tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		log.Fatalln("Error while opening file", err)
	}

	defer f.Close()

	if _, err = f.Write(jsonData); err != nil {
		log.Fatalln("Error while writing to JSON", err)
	}

}
