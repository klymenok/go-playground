package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	mydb "github.com/klymenok/go-playground/internal/db"
	"net/http"
	"strconv"
)

// GetComment godoc
// @Summary  get comment
// @Schemes
// @Description  get a comment by id
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Comment ID"  Format(int64)
// @Success      200  {object}  mydb.Comment
// @Router       /comments/{id} [get]
func GetComment(w http.ResponseWriter, r *http.Request) {
	commentId, _ := strconv.Atoi(chi.URLParam(r, "commentId"))
	comment, err := mydb.GetCommentById(int64(commentId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(comment)
	}
}

// GetComments godoc
// @Summary  get comment
// @Schemes
// @Description  get a comment by id
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Success      200  {array}  mydb.Comment
// @Router       /comments [get]
func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Comments"))
}

// CreateComment godoc
// @Summary  create comment
// @Schemes
// @Description  create comment
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        text        body      string  true  "Text"
// @Param        task        body      int     true  "Task"
// @Param        created_by  body      int     true  "Created By"
// @Success      201         {object}  mydb.Comment
// @Router       /comments [post]
func CreateComment(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	var comment mydb.Comment
	json.NewDecoder(r.Body).Decode(&comment)
	comment.Create()
	json.NewEncoder(w).Encode(comment)
}

// UpdateComment godoc
// @Summary  update comment
// @Schemes
// @Description  update an existing comment
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        id          path      int     true  "Comment ID"  Format(int64)
// @Param        text        body      string  true  "Text"
// @Param        task        body      int     true  "Task"
// @Param        created_by  body      int     true  "Created By"
// @Success      200         {object}  mydb.Task
// @Router       /comments/{id} [put]
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

// DeleteComment godoc
// @Summary  delete comment
// @Schemes
// @Description  delete an exising comment
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "Comment ID"  Format(int64)
// @Success      204
// @Router       /comments/{id} [delete]
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentId, _ := strconv.Atoi(chi.URLParam(r, "commentId"))
	mydb.DeleteCommentById(int64(commentId))
	w.Write([]byte("Comment deleted"))
}
