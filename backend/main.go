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
	r.HandleFunc("/tasks", getListTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", getOneTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", addTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id:[0-9]+}", updateTaskHandler).Methods("PATCH")

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
func getListTasksHandler(w http.ResponseWriter, r *http.Request) {
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
		// sqlの結果をtaskの各パラメータに格納する
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

func getOneTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	row := db.QueryRow("SELECT * FROM tasks WHERE ID=?", vars["id"])

	var task Task

	err := row.Scan(&task.Id, &task.Name, &task.Status)
	if err != nil {
		http.Error(w, "Scan Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// 挿入する値が正しいかどうか確認
	log.Printf("Received task: %+v", task)

	// データベースにタスクを挿入
	result, err := db.Exec("INSERT INTO tasks (name, status) VALUES (?, ?)", &task.Name, &task.Status)
	if err != nil {
		// エラーの詳細をログに出力
		log.Printf("Failed to insert task into database: %v", err)
		http.Error(w, "Failed to insert task into database", http.StatusInternalServerError)
		return
	}

	// 成功応答
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Task added successfully")

	// 挿入されたタスクのIDを取得し、ログに出力
	id, _ := result.LastInsertId()
	log.Printf("Inserted task ID: %d", id)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// 挿入する値が正しいかどうか確認
	log.Printf("Received task: %+v", task)

	// タスクを更新
	result, err := db.Exec("UPDATE tasks SET name = ? WHERE id = ? ", &task.Name, vars["id"])
	if err != nil {
		log.Printf("Failed to update task in database: %v", err)
		http.Error(w, "Failed to insert task into database", http.StatusInternalServerError)
		return
	}

	// 成功応答
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task updated successfully")

	// 更新されたタスクのIDを取得し、ログに出力
	id, _ := result.RowsAffected()
	log.Printf("Updated task ID: %d", id)
}
