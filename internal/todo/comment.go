package todo

import (
	"fmt"

	db "github.com/klymenok/go-playground/internal/db"
	"github.com/klymenok/go-playground/internal/utils"
)

type Comment struct {
	db *db.DB

	Id        int64  `json:"id"`
	Text      string `json:"text"`
	Task      int64  `json:"task"`
	CreatedBy int64  `json:"created_by"`
	CreatedAt string `json:"created_at"`
}

func NewComment(db *db.DB) *Comment {
	comment := &Comment{}
	comment.db = db
	return comment
}

func (c *Comment) Create() {
	// db := Connection()
	// defer db.Close()
	createCommentQuery := fmt.Sprintf(
		"insert into comment (text, task, created_by) values ('%s', %d, %d)",
		c.Text,
		c.Task,
		c.CreatedBy)
	result, err := c.db.Exec(createCommentQuery)
	utils.CheckError(err)
	c.Id, _ = result.LastInsertId()
}

func (c *Comment) Update() {
	// db := Connection()
	// defer db.Close()

	updateCommentQuery := fmt.Sprintf(
		"update comment set text='%s' where id=%d",
		c.Text,
		c.Id)
	_, err := c.db.Exec(updateCommentQuery)
	utils.CheckError(err)
}

func GetCommentById(commentId int64) (Comment, error) {
	var Comment Comment
	// db := Connection()
	// defer db.Close()

	// getCommentQuery := fmt.Sprintf("select * from comment where id=%d", commentId)
	// res := db.QueryRow(getCommentQuery)
	// err := res.Scan(&Comment.Id, &Comment.Text, &Comment.CreatedBy, &Comment.CreatedAt, &Comment.Task)
	// if err != nil {
	//   return Comment, err
	// }
	return Comment, nil
}

func DeleteCommentById(commentId int64) {
	// db := Connection()
	// defer db.Close()
	// deleteCommentQuery := fmt.Sprintf("delete from Comment where id=%d", commentId)
	// db.Exec(deleteCommentQuery)
}
