package todo

import (
	"fmt"

	"github.com/klymenok/go-playground/internal/db"
	"github.com/klymenok/go-playground/internal/utils"
)

type User struct {
	db        *db.DB
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewUser() *User {
	user := &User{}
	user.db = db.New()
	return user
}

func (u *User) Create() {
	createUserQuery := fmt.Sprintf(
		"insert into user (first_name, last_name) values ('%s', '%s')",
		u.FirstName,
		u.LastName)
	result, err := u.db.Exec(createUserQuery)
	utils.CheckError(err)
	u.Id, _ = result.LastInsertId()
}

func (u *User) Update() {

	updateUserQuery := fmt.Sprintf(
		"update user set first_name='%s', last_name='%s' where id=%d",
		u.FirstName,
		u.LastName,
		u.Id)
	_, err := u.db.Exec(updateUserQuery)
	utils.CheckError(err)
}
