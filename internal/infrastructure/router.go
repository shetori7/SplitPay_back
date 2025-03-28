package infrastructure

import (
	"SplitPay_back/config"
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "SplitPay_back/internal/interfaces/api"

	"github.com/gin-gonic/gin"
)

func Init() {
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

	e := gin.Default()
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	userController := controllers.NewUserController(NewSqlHandler(cfg))

	e.GET("/users", func(c *gin.Context) {
		users := userController.GetUser()
		c.Bind(&users)
		c.JSON(http.StatusOK, users)
	})

	//sample
	e.GET("/group/:id", func(c *gin.Context) {
		groupID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message":  "Group ID received",
			"group_id": groupID,
		})
	})

	//TODO：リクエストの構造体はここで持ってるでいいのか？
	type RequestBody struct {
		GroupID string   `json:"group_id"`
		Users   []string `json:"users"`
	}

	e.POST("/group/new", func(c *gin.Context) {
		var reqBody RequestBody

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "received array",
			"GroupID": reqBody.GroupID,
			"Users":   reqBody.Users,
		})
	})

	e.Run(":8000")
}
