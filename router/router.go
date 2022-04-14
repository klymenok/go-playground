package router

import (
	"github.com/go-chi/chi/v5"
	routehandlers "github.com/klymenok/go-playground/route-handlers"
	"net/http"
)

func Init() http.Handler {
	r := chi.NewRouter()
	r.Mount("/users/", userRouter())
	//r.Mount("tasks/", taskRouter())
	//r.Mount("comments/", commentRouter())
	return r
}

func userRouter() http.Handler {
	userRouter := chi.NewRouter()
	userRouter.Get("/", routehandlers.GetUsers)              // get all users
	userRouter.Get("/{userId}", routehandlers.GetUser)       // get user by id
	userRouter.Post("/", routehandlers.CreateUser)           // create new user
	userRouter.Put("/{userId}", routehandlers.UpdateUser)    // create new user
	userRouter.Delete("/{userId}", routehandlers.DeleteUser) // create new user
	return userRouter
}
