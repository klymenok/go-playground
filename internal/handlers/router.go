package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	userRouter.Get("/", GetUsers)              // get all users
	userRouter.Get("/{userId}", GetUser)       // get user by id
	userRouter.Post("/", CreateUser)           // create new user
	userRouter.Put("/{userId}", UpdateUser)    // update user
	userRouter.Delete("/{userId}", DeleteUser) // delete user

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
