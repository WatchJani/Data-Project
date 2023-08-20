package model

import (
	"encoding/json"

	e "root/check_error"

	"github.com/jmoiron/sqlx"
)

type User struct {
	name string
}

func GetAll(db *sqlx.DB) *[]User {
	var users []User

	db.Select(&users, "")

	return &users
}

func GetUser(db *sqlx.DB, userID string) ([]byte, error) {
	var user User

	e.ErrorHandler(db.Select(&user, "SELECT * FROM user WHERE id=?", userID))

	return json.Marshal(user)
}
