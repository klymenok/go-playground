package todo

import (
	"database/sql"
	"fmt"

	"github.com/klymenok/go-playground/internal/utils"
)

type User struct {
	db        *sql.DB
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *users {
	return &users{db}
}

func (u *users) Create(user *User) {
	createUserQuery := fmt.Sprintf(
		"insert into user (first_name, last_name) values ('%s', '%s')",
		user.FirstName,
		user.LastName)
	result, err := u.db.Exec(createUserQuery)
	utils.CheckError(err)
	user.Id, _ = result.LastInsertId()
}

func (u *users) Update(user User) {
	updateUserQuery := fmt.Sprintf(
		"update user set first_name='%s', last_name='%s' where id=%d",
		user.FirstName,
		user.LastName,
		user.Id)
	_, err := u.db.Exec(updateUserQuery)
	utils.CheckError(err)
}

func (u *users) ById(userId int64) (User, error) {
	var user User

	err := u.db.QueryRow(
		"select * from user where id=?", userId,
	).Scan(
		&user.Id, &user.FirstName, &user.LastName,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *users) DeleteById(userId int64) {
	deleteUserQuery := fmt.Sprintf("delete from user where id=%d", userId)
	u.db.Exec(deleteUserQuery)
}
