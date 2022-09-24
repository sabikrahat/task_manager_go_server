package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	allTasks()
	fmt.Println("Hello, World!")
	handleRoutes()
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", checkConnection).Methods("GET")
	//
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettasks/", getTasks).Methods("GET")
	//
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/gettask/{id}/", getTask).Methods("GET")
	//
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/create/", createTask).Methods("POST")
	//
	router.HandleFunc("/delete/{id}/", deleteTask).Methods("DELETE")
	router.HandleFunc("/delete/{id}/", deleteTask).Methods("DELETE")
	//
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/update/{id}/", updateTask).Methods("PUT")
	//
	log.Fatal(http.ListenAndServe(":1234", router))
}

func checkConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Backend is running...")
	fmt.Println("Endpoint Hit: checkConnection")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
	fmt.Println("Endpoint Hit: getTasks")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskId := mux.Vars(r)
	fmt.Println("getTask id: ", taskId)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == taskId["id"] {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"success": "false", "message": "Task not found"})
	}
	fmt.Println("Endpoint Hit: getTask")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createTask")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteTask")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateTask")
}

var tasks []Task

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func allTasks() {
	task := Task{ID: "1", Title: "Task One", Description: "This is a task", Date: "2020-01-01"}
	tasks = append(tasks, task)

	task = Task{ID: "2", Title: "Task Two", Description: "This is a task", Date: "2020-01-01"}
	tasks = append(tasks, task)

	task = Task{ID: "3", Title: "Task Three", Description: "This is a task", Date: "2020-01-01"}
	tasks = append(tasks, task)

	fmt.Println("All Tasks", tasks)
}
