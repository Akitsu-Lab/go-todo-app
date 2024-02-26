package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"log"
	"net/http"
	"os"
	"time"
)

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
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
	r.HandleFunc("/tasks", getListTasks)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
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

// List取得メソッド
func getListTasks(w http.ResponseWriter, r *http.Request) {
	// SELECT実行
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Taskに情報をセットする
	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Name, &task.Status)
		if err != nil {
			http.Error(w, "Scan Error", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}
	// JSON形式でレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
