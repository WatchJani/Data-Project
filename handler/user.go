package handler

import "github.com/jmoiron/sqlx"

type UserDB struct {
	db *sqlx.DB
}

func (db *UserDB) GetAll() {

}
