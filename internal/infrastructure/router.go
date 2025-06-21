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
	e.Use(cors.Default())

	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	//Controllerを作成
	userController := controllers.NewUserController(NewSqlHandler(cfg))
	groupController := controllers.NewGroupController(NewSqlHandler(cfg))
	paymentController := controllers.NewPaymentController(NewSqlHandler(cfg))

	// e.GET("/users", func(c *gin.Context) {
	// 	users := userController.GetUser()
	// 	c.Bind(&users)
	// 	c.JSON(http.StatusOK, users)
	// })

	e.GET("/users", func(c *gin.Context) {
		users := userController.GetUserByGroupId(c)
		c.Bind(&users)
		c.JSON(http.StatusOK, users)
	})

	e.GET("/group/getInfo", func(c *gin.Context) {
		group := groupController.GetGroupByUuId(c)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
			return
		}
		c.JSON(http.StatusOK, group)
	})

	e.POST("/group/new", func(c *gin.Context) {
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

	e.POST(("/payment/new"), func(c *gin.Context) {
		paymentController.Create(c)
	})

	e.Run(":8000")
}
