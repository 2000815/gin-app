package main

import (
	"fmt"
	"my-go-app/controller"
	"my-go-app/utility"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
    mode := os.Getenv("GIN_MODE")
    gin.SetMode(mode)
    router := controller.GetRouter()

    utility.Init()

    // サーバーを起動
    if err := router.Run(":8080"); err != nil {
        fmt.Printf("Failed to run server: %v\n", err)
    } else {
        fmt.Println("Server running on port 8080")
    }
}