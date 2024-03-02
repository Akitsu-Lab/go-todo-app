# API

## 環境変数

```shell
export DBUSER=root
export DBPASS=rootpass
export DBHOST=localhost
export DBPORT=3306
```

## 実行

```shell
go run main.go
```

## APIイメージ生成

```shell
docker build -t todo_backend .
```

## API

- 全権取得

```shell
curl http://localhost:8080/tasks | jq
```

- 1件取得

```shell
curl http://localhost:8080/tasks/2 | jq
```

- 追加
```shell
curl -X POST -H "Content-Type: application/json" -d '{"name":"追加したタスク", "status":1}' http://localhost:8080/tasks
```

- 更新
```shell
curl -X PATCH -H "Content-Type: application/json" -d '{"name":"更新したタスク"}' http://localhost:8080/tasks/1
```

- 削除
```shell
curl -X DELETE http://localhost:8080/tasks/1
```

## 参考

- [Golang(docker-hub)](https://hub.docker.com/_/mysql)
- [gorilla/mux](https://github.com/gorilla/mux)
