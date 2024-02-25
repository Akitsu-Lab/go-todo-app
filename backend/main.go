package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type task struct {
	Id     int
	Name   string
	Status int
}

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT"),
		DBName:               "todo",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to pint db", err)
	}

	log.Println("データベース接続完了")

}
