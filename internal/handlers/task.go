package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/klymenok/go-playground/internal/todo"
)

// @BasePath /api/v1

// GetTask godoc
// @Summary  get task
// @Schemes
// @Description  get a task by id
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"  Format(int64)
// @Success      200          {object}  todo.Task
// @Router       /tasks/{id} [get]
func GetTask(w http.ResponseWriter, r *http.Request) {
	todo := todo.NewToDo()

	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	task, err := todo.GetTaskById(int64(taskId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

// @BasePath /api/v1

// GetTasks godoc
// @Summary  get tasks
// @Schemes
// @Description  get list of all tasks
// @Tags         Task
// @Accept       json
// @Produce      json
// @Success      200  {array}  todo.Task
// @Router       /tasks [get]
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tasks"))
}

// @BasePath /api/v1

// CreateTask godoc
// @Summary  create a new task
// @Schemes
// @Description  create a new task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        title        body      string  true  "Title"
// @Param        description  body      string  true  "Description"
// @Param        created_by  body      int     true  "Created By"
// @Param        assignee     body      int     false  "Assignee"
// @Param        completed    body      bool    false  "Completed"
// @Success      201          {object}  todo.Task
// @Router       /tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	task := todo.NewTask()

	json.NewDecoder(r.Body).Decode(&task)
	task.Create()
	json.NewEncoder(w).Encode(task)
}

// @BasePath /api/v1

// UpdateTask godoc
// @Summary  update task
// @Schemes
// @Description  update an existing task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id           path      int     true  "Task ID"  Format(int64)
// @Param        title        body      string  true   "Title"
// @Param        description  body      string  true   "Description"
// @Param        created_by   body      int     true  "Created By"
// @Param        assignee     body      int     true  "Assignee"
// @Param        completed    body      bool    true  "Completed"
// @Success      200  {object}  todo.Task
// @Router       /tasks/{id} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	todo := todo.NewToDo()

	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	task, err := todo.GetTaskById(int64(taskId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewDecoder(r.Body).Decode(&task)
		task.Update()
		json.NewEncoder(w).Encode(task)
	}
}

// @BasePath /api/v1

// CompleteTask godoc
// @Summary  complete task
// @Schemes
// @Description  complete an existing task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"  Format(int64)
// @Success      200  {object}  todo.Task
// @Router       /tasks/{id}/complete [post]
func CompleteTask(w http.ResponseWriter, r *http.Request) {
	todo := todo.NewToDo()

	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	task, err := todo.GetTaskById(int64(taskId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		task.Complete()
		json.NewEncoder(w).Encode(task)
	}
}

// @BasePath /api/v1

// DeleteTask godoc
// @Summary  delete task
// @Schemes
// @Description  delete an existing task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "Task ID"  Format(int64)
// @Success      204
// @Router       /users/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	todo := todo.NewToDo()

	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	todo.DeleteTaskById(int64(taskId))
	w.Write([]byte("Task deleted"))
}

// CreateCommentForTask godoc
// @Summary  create comment for task
// @Schemes
// @Description  create a comment for the specified task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id          path      int     true  "Task ID"  Format(int64)
// @Param        text        body      string  true  "Text"
// @Param        created_by   body      int     true   "Created By"
// @Success      200         {object}  todo.Comment
// @Router       /tasks/{id}/create-comment [post]
func CreateCommentForTask(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	var comment todo.Comment
	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	json.NewDecoder(r.Body).Decode(&comment)
	comment.Task = int64(taskId)
	comment.Create()
	json.NewEncoder(w).Encode(comment)
}
