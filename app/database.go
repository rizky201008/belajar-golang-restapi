package app

import (
	"database/sql"
	"github.com/rizky201008/belajar-golang-restapi/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golang_restapi")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
