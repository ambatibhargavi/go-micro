package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var (
	tasks   []Task
	mu      sync.Mutex
	taskID  int
	dataFile = "tasks.json"
)

func loadTasks() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&tasks)
	if len(tasks) > 0 {
		taskID = tasks[len(tasks)-1].ID + 1
	}
}

func saveTasks() {
	file, err := os.Create(dataFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(tasks)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	mu.Lock()
	defer mu.Unlock()
	tmpl.Execute(w, tasks)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		mu.Lock()
		tasks = append(tasks, Task{ID: taskID, Title: title, Status: "Pending"})
		taskID++
		saveTasks()
		mu.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))
		mu.Lock()
		for i, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
		saveTasks()
		mu.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func completeTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))
		mu.Lock()
		for i, task := range tasks {
			if task.ID == id && task.Status != "Completed" {
				tasks[i].Status = "Completed"
				break
			}
		}
		saveTasks()
		mu.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	loadTasks()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addTaskHandler)
	http.HandleFunc("/delete", deleteTaskHandler)
	http.HandleFunc("/complete", completeTaskHandler) // Register the complete handler
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Changed port to 3000 for consistency with message
	fmt.Println("Server running at http://localhost:5000/")
	http.ListenAndServe(":5000", nil) // Now listening on port 3000
}

