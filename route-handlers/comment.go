package routehandlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	mydb "github.com/klymenok/go-playground/db"
	"net/http"
	"strconv"
)

func GetComment(w http.ResponseWriter, r *http.Request) {
	commentId, _ := strconv.Atoi(chi.URLParam(r, "commentId"))
	comment, err := mydb.GetCommentById(int64(commentId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(comment)
	}
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Comments"))
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	var comment mydb.Comment
	json.NewDecoder(r.Body).Decode(&comment)
	comment.Create()
	json.NewEncoder(w).Encode(comment)
}

func CreateCommentForTask(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	var comment mydb.Comment
	taskId, _ := strconv.Atoi(chi.URLParam(r, "taskId"))
	json.NewDecoder(r.Body).Decode(&comment)
	comment.Task = int64(taskId)
	comment.Create()
	json.NewEncoder(w).Encode(comment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	commentId, _ := strconv.Atoi(chi.URLParam(r, "commentId"))
	comment, err := mydb.GetCommentById(int64(commentId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewDecoder(r.Body).Decode(&comment)
		comment.Update()
		json.NewEncoder(w).Encode(comment)
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentId, _ := strconv.Atoi(chi.URLParam(r, "commentId"))
	mydb.DeleteCommentById(int64(commentId))
	w.Write([]byte("Comment deleted"))
}
