CREATE DATABASE IF NOT EXISTS quarkusdb;
USE quarkusdb;

-- ユーザーテーブル
CREATE TABLE IF NOT EXISTS users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL
);

-- 形式テーブル（3麻 or 4麻）
CREATE TABLE IF NOT EXISTS formats (
    format_id INT PRIMARY KEY,
    format_name VARCHAR(50) NOT NULL
);

-- 卓テーブル
CREATE TABLE IF NOT EXISTS tables (
    table_id INT PRIMARY KEY,
    is_reserved BOOLEAN NOT NULL,
    is_smoking BOOLEAN NOT NULL
);

-- 雀荘テーブル
CREATE TABLE IF NOT EXISTS mahjong_parlors (
    mahjong_parlor_id INT PRIMARY KEY,
    mahjong_parlor_name VARCHAR(255) NOT NULL,
    mahjong_parlor_address VARCHAR(765) NOT NULL
);

-- 予約テーブル
CREATE TABLE IF NOT EXISTS reservations (
    reservation_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    table_id INT NOT NULL,
    format_id INT NOT NULL,
    mahjong_parlor_id INT NOT NULL,
    temp_reserved BOOLEAN NOT NULL DEFAULT false,
    is_canceled BOOLEAN NOT NULL DEFAULT false,
    -- yyyy-mm-dd hh:mm:ss形式で保存
    start_date DATETIME NOT NULL,
    end_date DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (table_id) REFERENCES tables(table_id),
    FOREIGN KEY (format_id) REFERENCES formats(format_id),
    FOREIGN key (mahjong_parlor_id) REFERENCES mahjong_parlors(mahjong_parlor_id)
)CHARACTER SET utf8mb4;;

-- データの初期化
INSERT INTO users (user_name, email, password) VALUES ('Alice', 'alice@example.com', '123456');
INSERT INTO users (user_name, email, password) VALUES ('abe', 'abe@example.com', '123456');
INSERT INTO formats (format_id, format_name) VALUES (1, '3麻');
INSERT INTO formats (format_id, format_name) VALUES (2, '4麻');
INSERT INTO tables (table_id, is_reserved, is_smoking) VALUES (1, false,false);
INSERT INTO tables (table_id, is_reserved, is_smoking) VALUES (2, false,false);
INSERT INTO tables (table_id, is_reserved, is_smoking) VALUES (3, false,false);
INSERT INTO tables (table_id, is_reserved, is_smoking) VALUES (4, false,true);
INSERT INTO tables (table_id, is_reserved, is_smoking) VALUES (5, false,true);
INSERT INTO tables (table_id, is_reserved, is_smoking) VALUES (6, false,true);
INSERT INTO mahjong_parlors (mahjong_parlor_id, mahjong_parlor_name, mahjong_parlor_address) VALUES (1, '麻雀西部', '東京都新宿区高田馬場1-33-15');
INSERT INTO mahjong_parlors (mahjong_parlor_id, mahjong_parlor_name, mahjong_parlor_address) VALUES (2, '麻雀ZOO 梅田東通り店', '大阪府大阪市北区堂山町１−１４ こだまレジャービル 3F');
INSERT INTO reservations (user_id, table_id, format_id,mahjong_parlor_id, start_date, end_date) VALUES (1, 1, 1, 1,'2023-10-01 12:00:00', '2023-10-01 14:00:00');
INSERT INTO reservations (user_id, table_id, format_id, mahjong_parlor_id,temp_reserved,start_date, end_date) VALUES (1, 1, 1,1, true,'2023-10-01 12:00:00', '2023-10-01 14:00:00');
INSERT INTO reservations (user_id, table_id, format_id, mahjong_parlor_id,start_date, end_date) VALUES (2, 1, 1, 1,'2023-10-01 12:00:00', '2023-10-01 14:00:00');
INSERT INTO reservations (user_id, table_id, format_id, mahjong_parlor_id,start_date, end_date) VALUES (2, 1, 1, 2,'2023-10-01 12:00:00', '2023-10-01 14:00:00');
