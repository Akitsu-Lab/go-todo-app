# API

## 環境変数

```shell
export DBUSER=root
export DBPASS=rootpass
export DBHOST=localhost
export DBPORT=3306
```

## APIイメージ生成

```shell
docker build -t todo_backend .
```

## API

```shell
curl http://localhost:8080/tasks | jq
```

## 参考

- [Golang(docker-hub)](https://hub.docker.com/_/mysql)
- [gorilla/mux](https://github.com/gorilla/mux)
