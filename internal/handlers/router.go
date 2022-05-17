package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/klymenok/go-playground/internal/db"
	"github.com/klymenok/go-playground/internal/todo"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func Init() http.Handler {
	r := chi.NewRouter()
	manager := todo.NewManager(db.New())

	// middlewares
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// API docs
	r.Mount("/swagger", httpSwagger.WrapHandler)

	r.Mount("/users/", userRouter(manager))
	r.Mount("/tasks/", taskRouter())
	r.Mount("/comments/", commentRouter())
	return r
}

func initRestEndpoints(router chi.Router, handler Handler) {
	router.Get("/", handler.Get)
	router.Get("/{id}", handler.GetById)
	router.Post("/", handler.Create)
	router.Put("/{id}", handler.Update)
	router.Delete("/{id}", handler.Delete)
}

func userRouter(manager *todo.Manager) http.Handler {
	userRouter := chi.NewRouter()

	// main endpoints
	initRestEndpoints(userRouter, UserHandler{manager.Users})

	// additional endpoints
	userRouter.Post("/{userId}/create-task", CreateTaskForUser) // create new task for user
	return userRouter
}

func taskRouter() http.Handler {
	taskRouter := chi.NewRouter()

	// main endpoints
	taskRouter.Get("/", GetTasks)              // get all tasks
	taskRouter.Get("/{taskId}", GetTask)       // get task by id
	taskRouter.Post("/", CreateTask)           // create new task
	taskRouter.Put("/{taskId}", UpdateTask)    // update task
	taskRouter.Delete("/{taskId}", DeleteTask) // delete task

	// additional endpoints
	taskRouter.Post("/{taskId}/create-comment/", CreateCommentForTask) // create new comment for task
	taskRouter.Post("/{taskId}/complete/", CompleteTask)               // complete task

	return taskRouter
}

func commentRouter() http.Handler {
	commentRouter := chi.NewRouter()

	// main endpoints
	commentRouter.Get("/", GetComments)              // get all comments
	commentRouter.Get("/{userId}", GetComment)       // get comment by id
	commentRouter.Post("/", CreateComment)           // create new comment
	commentRouter.Put("/{userId}", UpdateComment)    // update comment
	commentRouter.Delete("/{userId}", DeleteComment) // delete comment

	return commentRouter
}
