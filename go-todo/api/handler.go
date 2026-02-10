package main

import(
	"encoding/json"
	"net/http"
	"go-todo/internal/tasks"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	list := tasks.LoadTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request){
	if r.Method !=http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var input struct {
		Task string `json:"task"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	taskList := tasks.LoadTasks()
	taskList = append(taskList, input.Task)
	tasks.SaveTasks(taskList)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task added",
	})
}


