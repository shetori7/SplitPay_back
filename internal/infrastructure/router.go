package infrastructure

import (
	"SplitPay_back/config"
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "SplitPay_back/internal/interfaces/api"

	"github.com/gin-contrib/cors"
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
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 必要に応じて本番用のドメインに変更
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//Controllerを作成
	userController := controllers.NewUserController(NewSqlHandler(cfg))
	groupController := controllers.NewGroupController(NewSqlHandler(cfg))
	paymentController := controllers.NewPaymentController(NewSqlHandler(cfg))

	// APIのルーティングを設定
	api := e.Group("/api")

	{
		// e.GET("/users", func(c *gin.Context) {
		// 	users := userController.GetUser()
		// 	c.Bind(&users)
		// 	c.JSON(http.StatusOK, users)
		// })

		api.GET("/users", func(c *gin.Context) {
			users := userController.GetUserByGroupId(c)
			c.Bind(&users)
			c.JSON(http.StatusOK, users)
		})

		api.GET("/group/getInfo", func(c *gin.Context) {
			group := groupController.GetGroupByUuId(c)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
				return
			}
			c.JSON(http.StatusOK, group)
		})

		api.POST("/group/new", func(c *gin.Context) {
			//TODO:リクエストとdomainのマッピング箇所をひとまとめにしたい
			g := groupController.Create(c)
			users := userController.CreateMultiple(c, g.GroupUuid)
			c.JSON(http.StatusOK, gin.H{
				"group_uuid": g.GroupUuid,
				"group_name": g.GroupName,
				"message":    "group created successfully",
				"Users":      users,
			})
		})

		api.POST(("/payment/new"), func(c *gin.Context) {
			paymentController.Create(c)
		})
	}

	e.Run(":" + cfg.ServerPort)
}
