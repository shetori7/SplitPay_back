# Waritomo開発

### 開発環境について
* Docker Desktopをインストールすること

### ローカルでのアプリケーション起動方法
```
# まずコンテナを立ち上げること（Nginxがたってないとアクセスできないため）

```

### Docker
* コンテナ起動手順（本プロジェクトではdocker-compose.ymlを使用して全コンテナを管理する）
```
my-quarkus-app/docker配下でdocker-composeを使用する
# コンテナの立ち上げ
docker compose up

# コンテナの停止
docker compose down

# コンテナ全削除
docker compose down --volumes --rmi all --remove-orphans

# エラーが出た場合
docker logs (コンテナ名→今回であればquarkus-mysql-container)
```
### MySQL
* MySQLへの接続
```
# MySQLに接続
docker exec -it waritomo-mysql-container mysql -u root -p

```
* 実行したSQLを確認したい場合

```
# dockerの中に入る
docker exec -it waritomo-mysql-container bash

# MySQL内で以下を実行
SET global general_log = 'ON';

# ログの出力先を確認
SHOW VARIABLES LIKE 'general_log_file';
tail -f /var/lib/mysql/$(hostname).log
```