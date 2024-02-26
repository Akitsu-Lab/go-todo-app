# go-todo-app

## 実行

環境変数を設定する
```shell
export DBUSER=root
export DBPASS=rootpass
export DBHOST=localhost
export DBPORT=3306
```

docker-compose立ち上げ

```shell
docker compose up -d
```

## 停止

```shell
docker compose down --rmi all --volumes --remove-orphans
```

## DBコンテナにMySQLログイン

- コンテナに入ってログイン
  ```shell
  docker exec -it todo_db /bin/bash
  ```
  ```shell
  mysql -u root -p
  ```
- コンテナに入らずにログイン
  ```shell
  mysql -u root -p --port=13306 --protocol=TCP
  ```

## 参考

- [Docker Composeの概要](https://matsuand.github.io/docs.docker.jp.onthefly/compose/)

