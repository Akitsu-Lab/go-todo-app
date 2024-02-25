package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"log"
	"net/http"
	"os"
	"time"
)

type Task struct {
	Id     int
	Name   string
	Status int
}

var db *sql.DB

func main() {
	// MySQL接続情報
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
	// Ping The Database
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping db", err)
	}

	log.Println("データベース接続完了")

	// DBからデータ取得

	// サーバーの起動
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "成功")
}
