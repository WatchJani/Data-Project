package model

import (
	"encoding/json"
	"time"

	e "root/check_error"
	"root/helper"

	"github.com/jmoiron/sqlx"

	"root/model/user/sql"
)

type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"lastname"`
	Password string    `json:"password"`
	Phone    string    `json:"phone"`
	Mail     string    `json:"mail"`
	UserName string    `json:"username"`
	BornDate time.Time `json:"borndate"`
	Created  time.Time `json:"created"`
}

func GetAll(db *sqlx.DB) ([]byte, error) {
	var users []User

	e.ErrorHandler(db.Select(&users, sql.GetAll()))

	return json.Marshal(users)
}

func GetUser(db *sqlx.DB, userID string) ([]byte, error) {
	var user User

	e.ErrorHandler(db.Get(&user, sql.GetUser(), userID))

	return json.Marshal(user)
}

func NewUser(db *sqlx.DB, Name, LastName, Password, Phone, Mail, UserName string, BornDate time.Time) {
	newPassword, err := helper.HashPassword(Password)
	e.ErrorHandler(err)

	_, err = db.Exec(sql.NewUser(), Name, LastName, newPassword, Phone, Mail, UserName, BornDate)
	e.ErrorHandler(err)

}
