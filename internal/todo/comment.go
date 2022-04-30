package todo

import (
	"fmt"

	"github.com/klymenok/go-playground/internal/db"
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

func NewComment() *Comment {
	comment := &Comment{}
	comment.db = db.New()
	return comment
}

func (c *Comment) Create() {
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

	updateCommentQuery := fmt.Sprintf(
		"update comment set text='%s' where id=%d",
		c.Text,
		c.Id)
	_, err := c.db.Exec(updateCommentQuery)
	utils.CheckError(err)
}
