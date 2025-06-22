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
### 本番環境について（http://150.230.195.18/）
* サーバー情報
```
WebサーバーとAPサーバーは同一サーバーでまかなう
OCI（オラクルクラウドの無料プランを使用している）

##WebAPサーバー
public IP：150.230.195.18
private IP：10.0.1.175
アカウント：walipay

##DBサーバー
public IP：なし
private IP：10.0.2.25
アカウント：walipay
その他DB設定については当プロジェクト内のprod.envを参照すること
```

* 接続方法
```
##WebAPサーバーへはSSHで接続を行うこと（秘密鍵の位置は適宜変えること）
ssh -i ~/.ssh/id_rsa opc@150.230.195.18

##DBサーバーへの接続
APサーバーでopcユーザーでログイン後/home/opc内に接続用スクリプトがある
```

* デプロイ方法
```
##フロントエンド
0.WebAPサーバーにSSHでログインする

1.ユーザーの切り替え
sudo su walipay

2.ディレクトリ移動し、シェルを実行する
cd /opt/walipay/src
./build_frontend.sh

3.ユーザーを切り替えてサービスを再起動
exit
sudo systemctl restart addpay_frontend.service

##バックエンド
0.WebAPサーバーにSSHでログインする

1.ユーザーの切り替え
sudo su walipay

2.ディレクトリ移動し、シェルを実行する
cd /opt/walipay/src
./build_backend.sh

3.ユーザーを切り替えてサービスを再起動
exit
sudo systemctl restart addpay_backend.service

##Nginx
Nginxの設定ファイルを編集した際はNginxサービスの再起動を行うこと
1.opcユーザーになる

2.Nginxの再起動
sudo systemctl restart nginx
```
