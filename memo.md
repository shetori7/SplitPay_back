# Waritomo開発

## 開発環境について
- Docker Desktopをインストールすること

## ローカルでのアプリケーション起動方法
```
# まずコンテナを立ち上げること（Nginxがたってないとアクセスできないため）

# アクセス方法
http://localhost:8080/{APIのパスを入力}
```

## Docker
- コンテナ起動手順（本プロジェクトではdocker-compose.ymlを使用して全コンテナを管理する）
```
my-quarkus-app/docker配下でdocker-composeを使用する
# コンテナの立ち上げ（docker desktop立ち上げてたほうがいいかも）
docker-compose up

# コンテナの停止
docker-compose down
```
### MySQL
- MySQLへの接続
```
# MySQLに接続（コンテナ起動後しばらくつながらないため注意、エラーになっても再度下記コマンドを実行することでつながる）
docker exec -it waritomo-mysql-container mysql -u root -p

# エラーが出た場合
docker logs (コンテナ名→今回であればquarkus-mysql-container)
```