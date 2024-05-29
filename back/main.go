package main

import (
	"Bookapp/book_methods"
	"Bookapp/admin_methods"
	"Bookapp/ent"
	"Bookapp/sign"
	"Bookapp/structs"
	"Bookapp/token"
	"context"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)




func main() {
	
	//Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()

	//PostgreSQLに接続
	client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=book_db password=password sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	//グローバル変数に代入
	structs.Client=client
	
	// CORS設定
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	//ルーター
	router.POST("users/sign_up", sign.SignUpHandler)//サインアップ
	router.POST("users/sign_in", sign.SignInHandler)//サインイン
	router.POST("users/Tokencheck", token.TokenCheckHandler)//トークンチェック
	router.POST("books/resgister",book_methods.RegisterHandler)//本登録
	router.POST("books/get_list",book_methods.GetListHandler)//ユーザーの本リスト取得
	router.POST("books/update",book_methods.UpdateHandler)//本のアップデート
	router.POST("books/delete",book_methods.DeleteHandler)//選択削除処理
	router.POST("admin/sign_in",admin_methods.AdminSignInHandler)//管理者サインイン
	router.POST("admin/book_resgister",admin_methods.SuggestRegisterHandler)//おすすめ本登録
	router.POST("admin/book_get_list",admin_methods.SuggestGetListHandler)//おすすめ本リスト取得
	router.POST("admin/book_update",admin_methods.SuggestUpdateHandler)//おすすめ本アップデート
	router.POST("admin/book_delete",admin_methods.SuggestDeleteHandler)//おすすめ本選択削除処理
	// サーバー起動
	router.Run(":8080")
}
