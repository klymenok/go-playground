package mydb

import "log"

type Comment struct {
	Id        int64  `json:"id"`
	Text      string `json:"text"`
	Task      int64  `json:"task"`
	CreatedBy int64  `json:"created_by"`
	CreatedAt string `json:"created_at"`
}

func (c *Comment) Create() {
	db := Connection()
	defer db.Close()
	createCommentQuery := `insert into comment (text, task, created_by) values ('$1', $2, $3)`
	result, err := db.Exec(createCommentQuery, c.Text, c.Task, c.CreatedBy)
	if err != nil {
		log.Fatalln(err)
	}
	c.Id, _ = result.LastInsertId()
}

func (c *Comment) Update() {
	db := Connection()
	defer db.Close()

	updateCommentQuery := `update comment set text='$1' where id=$2`
	_, err := db.Exec(updateCommentQuery, c.Text, c.Id)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetCommentById(commentId int64) (Comment, error) {
	var Comment Comment
	db := Connection()
	defer db.Close()

	getCommentQuery := `select * from comment where id=$1`
	res := db.QueryRow(getCommentQuery, commentId)
	err := res.Scan(&Comment.Id, &Comment.Text, &Comment.CreatedBy, &Comment.CreatedAt, &Comment.Task)
	if err != nil {
		return Comment, err
	}
	return Comment, nil
}

func DeleteCommentById(commentId int64) {
	db := Connection()
	defer db.Close()
	deleteCommentQuery := `delete from Comment where id=$1`
	db.Exec(deleteCommentQuery, commentId)
}
