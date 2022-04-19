package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	routehandlers "github.com/klymenok/go-playground/route-handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func Init() http.Handler {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// API docs
	r.Mount("/swagger", httpSwagger.WrapHandler)

	r.Mount("/users/", userRouter())
	r.Mount("/tasks/", taskRouter())
	r.Mount("/comments/", commentRouter())
	return r
}

func userRouter() http.Handler {
	userRouter := chi.NewRouter()

	// main endpoints
	userRouter.Get("/", routehandlers.GetUsers)              // get all users
	userRouter.Get("/{userId}", routehandlers.GetUser)       // get user by id
	userRouter.Post("/", routehandlers.CreateUser)           // create new user
	userRouter.Put("/{userId}", routehandlers.UpdateUser)    // update user
	userRouter.Delete("/{userId}", routehandlers.DeleteUser) // delete user

	// additional endpoints
	userRouter.Post("/{userId}/create-task", routehandlers.CreateTaskForUser) // create new task for user
	return userRouter
}

func taskRouter() http.Handler {
	taskRouter := chi.NewRouter()

	// main endpoints
	taskRouter.Get("/", routehandlers.GetTasks)              // get all tasks
	taskRouter.Get("/{taskId}", routehandlers.GetTask)       // get task by id
	taskRouter.Post("/", routehandlers.CreateTask)           // create new task
	taskRouter.Put("/{taskId}", routehandlers.UpdateTask)    // update task
	taskRouter.Delete("/{taskId}", routehandlers.DeleteTask) // delete task

	// additional endpoints
	taskRouter.Post("/{taskId}/create-comment/", routehandlers.CreateCommentForTask) // create new comment for task
	taskRouter.Post("/{taskId}/complete/", routehandlers.CompleteTask)               // complete task

	return taskRouter
}

func commentRouter() http.Handler {
	commentRouter := chi.NewRouter()

	// main endpoints
	commentRouter.Get("/", routehandlers.GetComments)              // get all comments
	commentRouter.Get("/{userId}", routehandlers.GetComment)       // get comment by id
	commentRouter.Post("/", routehandlers.CreateComment)           // create new comment
	commentRouter.Put("/{userId}", routehandlers.UpdateComment)    // update comment
	commentRouter.Delete("/{userId}", routehandlers.DeleteComment) // delete comment

	return commentRouter
}
