package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	*sqlx.DB
}

type Configuration struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// By .env
//
// DB_HOST DB_PORT DB_USER DB_PASSWORD DB_NAME
func PostgresConfiguration() *Configuration {
	return &Configuration{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}

func CustomPostgresConfiguration(Host, Port, User, Password, Name string) *Configuration {
	return &Configuration{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		Name:     Name,
	}
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db}
}

func Open(config *Configuration) (*Store, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	store := NewStore(db)
	return store, nil
}

func (store *Store) CheckStoreConnection() error {
	return store.Ping()
}

func (store *Store) Close() error {
	return store.Close()
}
