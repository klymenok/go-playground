package todo

import (
	"database/sql"
)

type Manager struct {
	Tasks    *tasks
	Users    UserManager
	Comments *comments
}

func NewManager(db *sql.DB) *Manager {
	manager := &Manager{}
	manager.Tasks = NewTask(db)
	manager.Users = NewUsers(db)
	manager.Comments = NewComments(db)
	return manager
}
