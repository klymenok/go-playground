package mydb

import (
	"log"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *User) Create() {
	db := Connection()
	defer db.Close()
	createUserQuery := `insert into user (first_name, last_name) values ('$1', '$2')`
	result, err := db.Exec(createUserQuery, u.FirstName, u.LastName)
	if err != nil {
		log.Fatalln(err)
	}
	u.Id, _ = result.LastInsertId()
}

func (u *User) Update() {
	db := Connection()
	defer db.Close()

	updateUserQuery := `update user set first_name='$1', last_name='$2' where id=$3`
	_, err := db.Exec(updateUserQuery, u.FirstName, u.LastName, u.Id)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetUserById(userId int64) (User, error) {
	var user User
	db := Connection()
	defer db.Close()

	getUserQuery := `select * from user where id=$1`
	res := db.QueryRow(getUserQuery, userId)
	err := res.Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		return user, err
	}
	return user, nil
}

func DeleteUserById(userId int64) {
	db := Connection()
	defer db.Close()
	deleteUserQuery := `delete from user where id=$1`
	db.Exec(deleteUserQuery, userId)
}
