package todo

import (
	"fmt"

	"github.com/klymenok/go-playground/internal/db"
)

type ToDo struct {
	db *db.DB
}

func NewToDo() *ToDo {
	todo := &ToDo{}
	todo.db = db.New()
	return todo
}

func (t *ToDo) GetTaskById(taskId int64) (Task, error) {
	var task Task

	getTaskQuery := fmt.Sprintf("select * from task where id=%d", taskId)
	res := t.db.QueryRow(getTaskQuery)
	err := res.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedBy, &task.Assignee, &task.Completed)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (t *ToDo) DeleteTaskById(taskId int64) {
	deleteTaskQuery := fmt.Sprintf("delete from task where id=%d", taskId)
	t.db.Exec(deleteTaskQuery)
}

func (t *ToDo) GetUserById(userId int64) (User, error) {
	var user User

	getUserQuery := fmt.Sprintf("select * from user where id=%d", userId)
	res := t.db.QueryRow(getUserQuery)
	err := res.Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (t *ToDo) DeleteUserById(userId int64) {
	deleteUserQuery := fmt.Sprintf("delete from user where id=%d", userId)
	t.db.Exec(deleteUserQuery)
}

func (t *ToDo) GetCommentById(commentId int64) (Comment, error) {
	var Comment Comment

	getCommentQuery := fmt.Sprintf("select * from comment where id=%d", commentId)
	res := t.db.QueryRow(getCommentQuery)
	err := res.Scan(&Comment.Id, &Comment.Text, &Comment.CreatedBy, &Comment.CreatedAt, &Comment.Task)
	if err != nil {
		return Comment, err
	}
	return Comment, nil
}

func (t *ToDo) DeleteCommentById(commentId int64) {
	deleteCommentQuery := fmt.Sprintf("delete from Comment where id=%d", commentId)
	t.db.Exec(deleteCommentQuery)
}
