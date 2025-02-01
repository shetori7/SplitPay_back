# ベースイメージとしてMySQLを指定
FROM mysql:latest

# 環境変数でMySQLの設定を指定
ENV MYSQL_ROOT_PASSWORD=root

# 初期SQLファイルをコピー
COPY ./init.sql /docker-entrypoint-initdb.d/

# ホストのmy.cnfファイルをコンテナにコピー
COPY my.cnf /etc/mysql/my.cnf

# ファイルのアクセス権を設定
RUN chmod 644 /etc/mysql/my.cnf && chown root:root /etc/mysql/my.cnf

# デフォルトポート3306を公開
EXPOSE 3306
