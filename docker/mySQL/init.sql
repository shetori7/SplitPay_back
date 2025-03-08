CREATE DATABASE IF NOT EXISTS waritomodb;
USE waritomodb;

-- グループテーブル
CREATE TABLE IF NOT EXISTS wari_groups (
    group_id INT AUTO_INCREMENT PRIMARY KEY,
    group_name VARCHAR(50) NOT NULL
);

-- ユーザーテーブル
CREATE TABLE IF NOT EXISTS wari_users (
    user_id INT AUTO_INCREMENT,
    group_id INT,
    user_name VARCHAR(50) NOT NULL,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (group_id) REFERENCES wari_groups(group_id)
);

-- データの初期化
INSERT INTO wari_groups (group_name) VALUES ('group1');
INSERT INTO wari_users (group_id, user_name) VALUES (1, 'user1');