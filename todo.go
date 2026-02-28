package main

import(
	"encoding/json"
	"os"
	"time"
)

const filename = "tasks.json"

type Todo struct{
	Number int `json:"number"`
	Title string `json:"title"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	CompletedAt time.Time `json:"completedAt"`
}

func Edit(id int, title string) error{
	var tasks []Todo

	data,err := os.ReadFile("tasks.json")

	if err == nil {
        json.Unmarshal(data, &tasks)
    }

	tasks[id-1].Title = title

	updatedData,err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }

	err = os.WriteFile(filename, updatedData, 0644)
    if err != nil {
        return err
    }

	return nil
}

func Toggle(id int) error{
	var tasks []Todo

	data,err := os.ReadFile("tasks.json")

	if err == nil {
        json.Unmarshal(data, &tasks)
    }

	tasks[id-1].Status = "done"
	tasks[id-1].CompletedAt = time.Now()

	updatedData,err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }

	err = os.WriteFile(filename, updatedData, 0644)
    if err != nil {
        return err
    }

	return nil
}

func Read() ([]Todo, error){
	
	var tasks []Todo

	data, err := os.ReadFile(filename)

	json.Unmarshal(data,&tasks)

	return tasks, err
}

func Write(task Todo) error{
	var tasks []Todo

	data,err := os.ReadFile("tasks.json")

	if err == nil {
        json.Unmarshal(data, &tasks)
    }

	task.Number = len(tasks) + 1

	tasks = append(tasks,task)

	updatedData,err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }

	err = os.WriteFile(filename, updatedData, 0644)
    if err != nil {
        return err
    }

    return nil
}