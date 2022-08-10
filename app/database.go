package app

import (
	"database/sql"
	"jeryhardianto/golang-tugas/helper"
	"time"
)

func NewDB() *sql.DB {
	//db, err := sql.Open("mysql", "root:123456789@tcp(localhost:3306)/golang_tugas")
	//Host MySQL HeroKu
	db, err := sql.Open("mysql", "bbd3a79635f3f8:eb4c44e6@tcp(us-cdbr-east-06.cleardb.net:3306)/golang_tugas")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
