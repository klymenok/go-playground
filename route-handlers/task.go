package routehandlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	mydb "github.com/klymenok/go-playground/db"
	"net/http"
	"strconv"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	task, err := mydb.GetTaskById(int64(taskId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tasks"))
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	var task mydb.Task
	json.NewDecoder(r.Body).Decode(&task)
	task.Create()
	json.NewEncoder(w).Encode(task)
}

func CreateTaskForUser(w http.ResponseWriter, r *http.Request) {
	var task mydb.Task
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	json.NewDecoder(r.Body).Decode(&task)
	task.CreatedBy = int64(userId)
	task.Create()
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	task, err := mydb.GetTaskById(int64(taskId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewDecoder(r.Body).Decode(&task)
		task.Update()
		json.NewEncoder(w).Encode(task)
	}
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	task, err := mydb.GetTaskById(int64(taskId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		task.Complete()
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	mydb.DeleteTaskById(int64(taskId))
	w.Write([]byte("Task deleted"))
}
