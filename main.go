package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"SplitPay_back/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// コマンドライン引数を取得
	// 最初の引数はプロファイル名、デフォルトは "dev" とする
	profile := "localhost" // デフォルトプロファイル
	if len(os.Args) > 1 {
		profile = os.Args[1] // コマンドライン引数が指定されていれば、それをプロファイルとして使う
	}

	// 設定を読み込む
	cfg, err := config.LoadConfig(profile)
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// 設定を表示
	fmt.Println("Loaded Config:")
	fmt.Println("Database Host:", cfg.DBHost)
	fmt.Println("Server Port:", cfg.ServerPort)

	// データベース接続設定（GORM）
	// db, err := infrastructure.ConnectDB(cfg)
	// if err != nil {
	// 	log.Fatal("Error connecting to database: ", err)
	// }

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":8000")
}

// package main

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	// gorm.Modelをつけると、idとCreatedAtとUpdatedAtとDeletedAtが作られる
// 	gorm.Model

// 	Name     string
// 	Age      int
// 	IsActive bool
// }

// func main() {
// 	// dbを作成します
// 	db := dbInit()

// 	// dbをmigrateします
// 	db.AutoMigrate(&User{})
// }

// func dbInit() *gorm.DB {
// 	dsn := "root:root@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=true"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	return db
// }
