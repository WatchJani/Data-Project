package model

type User struct {
	name string
}

func GetAll() *[]User {
	var users []User

	return &users
}
