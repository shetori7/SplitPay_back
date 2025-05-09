CREATE DATABASE IF NOT EXISTS waritomodb;
USE waritomodb;

-- グループテーブル
CREATE TABLE IF NOT EXISTS wari_groups (
    group_id INT AUTO_INCREMENT PRIMARY KEY,
    group_name VARCHAR(50) NOT NULL,
    group_uuid VARCHAR(50) NOT NULL UNIQUE,
    group_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    group_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ユーザーテーブル
CREATE TABLE IF NOT EXISTS wari_users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    group_id INT,
    user_name VARCHAR(50) NOT NULL,
    FOREIGN KEY (group_id) REFERENCES wari_groups(group_id),
    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    FOREIGN KEY (wari_loans_from_user_id) REFERENCES wari_users(user_id)
    wari_loans_amount DECIMAL(10, 2) NOT NULL,
);

-- 支払いのテーブル（誰がいくら払ったか）
CREATE TABLE IF NOT EXISTS wari_payments (
    payment_id INT AUTO_INCREMENT PRIMARY KEY,
    group_id INT,
    user_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    message VARCHAR(255),
    FOREIGN KEY (group_id) REFERENCES wari_groups(group_id),
    FOREIGN KEY (user_id) REFERENCES wari_users(user_id)
);

-- 立替してもらった人のテーブル（誰がいくら立替えてもらったか）
CREATE TABLE IF NOT EXISTS wari_loans (
    loan_id INT AUTO_INCREMENT PRIMARY KEY,
    payment_id INT,
    group_id INT,
    to_user_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (group_id) REFERENCES wari_groups(group_id),
    FOREIGN KEY (payment_id) REFERENCES wari_payments(payment_id),
    FOREIGN KEY (to_user_id) REFERENCES wari_users(user_id),
);

-- 最終的な支払いのテーブル（誰がいくら払ったか、全方向に向きをもつ、正と負で向きを判断する）
-- UPDATEで更新する
CREATE TABLE IF NOT EXISTS wari_final_payments (
    final_payment_id INT AUTO_INCREMENT PRIMARY KEY,
    group_id INT,
    from_user_id INT,
    to_user_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (group_id) REFERENCES wari_groups(group_id),
    FOREIGN KEY (from_user_id) REFERENCES wari_users(user_id),
    FOREIGN KEY (to_user_id) REFERENCES wari_users(user_id),
);

-- データの初期化
INSERT INTO wari_groups (group_name,group_uuid) VALUES ('group1', 'uuid1');
INSERT INTO wari_users (group_id, user_name) VALUES (1, 'user1');