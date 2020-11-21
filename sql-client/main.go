package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbClient *sql.DB

const queryGetUser = "SELECT id, email FROM users WHERE id=%d"

type User struct {
	Id    int64
	Email string
}

func init() {
	var err error
	dbClient, err = sql.Open("mysql", "path")
	if err != nil {
		panic(err)
	}
}

func main() {
	user, err := GetUser(123)
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
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user does not exist")
}
