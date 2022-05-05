package todo

import (
	"database/sql"
	"fmt"

	"github.com/klymenok/go-playground/internal/utils"
)

type Comment struct {
	Id        int64  `json:"id"`
	Text      string `json:"text"`
	Task      int64  `json:"task"`
	CreatedBy int64  `json:"created_by"`
	CreatedAt string `json:"created_at"`
}

type comments struct {
	db *sql.DB
}

func NewComments(db *sql.DB) *comments {
	return &comments{db}
}

func (c *comments) Create(comment *Comment) {
	createCommentQuery := fmt.Sprintf(
		"insert into comment (text, task, created_by) values ('%s', %d, %d)",
		comment.Text,
		comment.Task,
		comment.CreatedBy)
	result, err := c.db.Exec(createCommentQuery)
	utils.CheckError(err)
	comment.Id, _ = result.LastInsertId()
}

func (c *comments) Update(comment Comment) {

	updateCommentQuery := fmt.Sprintf(
		"update comment set text='%s' where id=%d",
		comment.Text,
		comment.Id)
	_, err := c.db.Exec(updateCommentQuery)
	utils.CheckError(err)
}

func (c *comments) GetById(commentId int64) (Comment, error) {
	var Comment Comment

	getCommentQuery := fmt.Sprintf("select * from comment where id=%d", commentId)
	res := c.db.QueryRow(getCommentQuery)
	err := res.Scan(&Comment.Id, &Comment.Text, &Comment.CreatedBy, &Comment.CreatedAt, &Comment.Task)
	if err != nil {
		return Comment, err
	}
	return Comment, nil
}

func (c *comments) DeleteById(commentId int64) {
	deleteCommentQuery := fmt.Sprintf("delete from Comment where id=%d", commentId)
	c.db.Exec(deleteCommentQuery)
}
