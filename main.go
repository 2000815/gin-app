package main

import "github.com/gin-gonic/gin"

func main() {
    // Gin Routerを作成
    router := gin.Default()

    // ハンドラー関数を登録
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello, World!"})
    })

    // サーバーを起動
    router.Run(":8080")
}