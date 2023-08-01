package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
}

type allTasks []Task

var tasks = allTasks{
	{
		Id:        1,
		Name:      "Task one",
		Desc:      "Wash dishes",
		Completed: false,
	},
	{
		Id:        2,
		Name:      "Task Two",
		Desc:      "Do Homework",
		Completed: false,
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", welcome)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")

	log.Println("Server running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Endpoints
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Client!")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal("The request is not valid")
	}

	json.Unmarshal(body, &newTask)
	newTask.Id = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTask, _ := strconv.Atoi(vars["id"])

	for _, task := range tasks {
		if task.Id == idTask {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTask, _ := strconv.Atoi(vars["id"])

	for i, task := range tasks {
		if task.Id == idTask {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprint(w, "The task has been deleted successfully")
		}
	}
}
