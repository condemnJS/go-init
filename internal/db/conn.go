package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Database interface {
	Init() *sqlx.DB
}

type DB struct {
	Driver   string
	User     string
	Password string
	DBName   string
	SslMode  string
}

func NewDB() *DB {
	return &DB{
		"mysql",
		"root",
		"root",
		"goolang",
		"disable",
	}
}

func (db *DB) Init() *sqlx.DB {
	connRules := fmt.Sprintf("%s:%s@/%s", db.User, db.Password, db.DBName)

	conn, err := sqlx.Connect(db.Driver, connRules)

	if err != nil {
		log.Fatalln(err)
	}

	return conn
}
