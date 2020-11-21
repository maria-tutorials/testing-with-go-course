package main

import (
	"errors"
	"fmt"

	"./sqlclient"
)

var dbClient sqlclient.SqlClient //*sql.DB

const queryGetUser = "SELECT id, email FROM users WHERE id=?;"

type User struct {
	Id    int64
	Email string
}

func init() {
	var err error
	dbClient, err = sqlclient.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "password", "127.0.0.1:3306", "users_db"))
	if err != nil {
		panic(err)
	}
}

func main() {
	user, err := GetUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Email)
}

func GetUser(id int64) (*User, error) {
	rows, err := dbClient.Query(queryGetUser, id)
	if err != nil {
		return nil, err
	}
	user := User{}
	for rows.HasNext() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user does not exist")
}
