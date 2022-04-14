package routehandlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	mydb "github.com/klymenok/go-playground/db"
	"log"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	user, err := mydb.GetUserById(int64(userId))
	log.Println(err)
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	var user mydb.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Create()
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	user, err := mydb.GetUserById(int64(userId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewDecoder(r.Body).Decode(&user)
		user.Update()
		json.NewEncoder(w).Encode(user)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	mydb.DeleteUserById(int64(userId))
	w.Write([]byte("User deleted"))
}
