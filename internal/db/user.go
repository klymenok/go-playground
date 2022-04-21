package mydb

import (
	"fmt"
	"github.com/klymenok/go-playground/internal/utils"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *User) Create() {
	db := Connection()
	defer db.Close()
	createUserQuery := fmt.Sprintf(
		"insert into user (first_name, last_name) values ('%s', '%s')",
		u.FirstName,
		u.LastName)
	result, err := db.Exec(createUserQuery)
	utils.CheckError(err)
	u.Id, _ = result.LastInsertId()
}

func (u *User) Update() {
	db := Connection()
	defer db.Close()

	updateUserQuery := fmt.Sprintf(
		"update user set first_name='%s', last_name='%s' where id=%d",
		u.FirstName,
		u.LastName,
		u.Id)
	_, err := db.Exec(updateUserQuery)
	utils.CheckError(err)
}

func GetUserById(userId int64) (User, error) {
	var user User
	db := Connection()
	defer db.Close()

	getUserQuery := fmt.Sprintf("select * from user where id=%d", userId)
	res := db.QueryRow(getUserQuery)
	err := res.Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		return user, err
	}
	return user, nil
}

func DeleteUserById(userId int64) {
	db := Connection()
	defer db.Close()
	deleteUserQuery := fmt.Sprintf("delete from user where id=%d", userId)
	db.Exec(deleteUserQuery, userId)
}
