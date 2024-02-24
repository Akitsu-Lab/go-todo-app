# go-todo-app

## 実行

1. DBイメージ生成

   ```shell
   docker build -t todo_db ./db
   ```

2. APIイメージ生成 
   ```shell
   🚧 工事中
   ```

3. UIイメージ生成
   ```shell
   🚧 工事中
   ```
4. docker-compose立ち上げ
   ```shell
   docker-compose up -d
   ```
## 停止
```shell
docker-compose down
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

