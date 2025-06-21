CREATE DATABASE IF NOT EXISTS waritomodb;
USE waritomodb;

-- グループテーブル
CREATE TABLE IF NOT EXISTS wari_groups (
    group_name VARCHAR(50) NOT NULL,
    group_uuid VARCHAR(50) NOT NULL PRIMARY KEY,
    group_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    group_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ユーザーテーブル
CREATE TABLE IF NOT EXISTS wari_users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    group_uuid VARCHAR(50) NOT NULL,
    user_name VARCHAR(50) NOT NULL,
    FOREIGN KEY (group_uuid) REFERENCES wari_groups(group_uuid),
    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 立替者の情報と支払額を保持するテーブル
CREATE TABLE IF NOT EXISTS wari_payments (
    payment_id INT AUTO_INCREMENT PRIMARY KEY,
    payer_group_id VARCHAR(50) NOT NULL,
    payer_user_id INT,
    payer_amount DECIMAL(10, 2) NOT NULL,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    message VARCHAR(255),
    FOREIGN KEY (payer_group_id) REFERENCES wari_groups(group_uuid),
    FOREIGN KEY (payer_user_id) REFERENCES wari_users(user_id)
);

-- 立替してもらった人のテーブル
CREATE TABLE IF NOT EXISTS wari_loans (
    loan_id INT AUTO_INCREMENT PRIMARY KEY,
    payment_id INT,
    payee_amount DECIMAL(10, 2) NOT NULL,
    payee_user_id INT,
    FOREIGN KEY (payment_id) REFERENCES wari_payments(payment_id),
    FOREIGN KEY (payee_user_id) REFERENCES wari_users(user_id)
);

-- 最終的な支払いのテーブル（誰がいくら払ったか、全方向に向きをもつ、正と負で向きを判断する）
-- UPDATEで更新する
CREATE TABLE IF NOT EXISTS wari_final_payments (
    final_payment_id INT AUTO_INCREMENT PRIMARY KEY,
    group_uuid VARCHAR(50) NOT NULL,
    from_user_id INT,
    to_user_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (group_uuid) REFERENCES wari_groups(group_uuid),
    FOREIGN KEY (from_user_id) REFERENCES wari_users(user_id),
    FOREIGN KEY (to_user_id) REFERENCES wari_users(user_id)
);
