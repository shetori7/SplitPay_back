// // DB接続
// package infrastructure

// import (
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // ConnectDB はデータベースに接続する
// func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
// 	// 接続文字列を作成
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

// 	// DB接続
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Error connecting to database: %v", err)
// 		return nil, err
// 	}

// 	return db, nil
// }