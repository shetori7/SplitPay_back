// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

//	func main() {
//		engine := gin.Default()
//		engine.GET("/", func(c *gin.Context) {
//			c.JSON(http.StatusOK, gin.H{
//				"message": "hello world",
//			})
//		})
//		engine.Run(":8000")
//	}
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Modelをつけると、idとCreatedAtとUpdatedAtとDeletedAtが作られる
	gorm.Model

	Name     string
	Age      int
	IsActive bool
}

func main() {
	// dbを作成します
	db := dbInit()

	// dbをmigrateします
	db.AutoMigrate(&User{})
}

func dbInit() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
