package todo

import (
	"database/sql"
)

type ToDo struct {
	Tasks    *tasks
	Users    *users
	Comments *comments
}

func NewToDo(db *sql.DB) *ToDo {
	todo := &ToDo{}
	todo.Tasks = NewTask(db)
	todo.Users = NewUsers(db)
	todo.Comments = NewComments(db)
	return todo
}
