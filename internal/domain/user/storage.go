package user

import "github.com/jmoiron/sqlx"

type RepositoryInterface interface {
	Create()
	First() *User
}

type Storage struct {
	DB *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db}
}

func (storage *Storage) Create() {
	storage.DB.MustExec("INSERT INTO users (email, password) VALUES ($1, $2)", "tatarin@mail.ru", "tatarin")
}

func (storage *Storage) First() *User {
	//TODO implement me
	panic("implement me")
}
