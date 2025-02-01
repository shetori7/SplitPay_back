# nginxのDockerfile
FROM nginx:1.19.6

# カスタムのnginx.confを追加
COPY default.conf /etc/nginx/conf.d/

# Nginxを起動
CMD ["nginx", "-g", "daemon off;"]