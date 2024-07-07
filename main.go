package main

import (
	"fmt"
	"my-go-app/controller"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// func main() {
//     // fmt.Println("こちらテストです")

//     // // データベース初期

//     // router := controller.GetRouter()

//     // // サーバーを起動
//     // if err := router.Run(":8080"); err != nil {
//     //     fmt.Printf("Failed to run server: %v\n", err)
//     // }

//     router := gin.Default()
//     router.GET("/", func(c *gin.Context) {
// 		c.HTML(200, "index.html", gin.H{
// 			"title": "Welcome!",
// 		})
// 	})
// }


func main() {
    fmt.Println("こちらテストです")

    router := controller.GetRouter()

    // サーバーを起動
    if err := router.Run(":8080"); err != nil {
        fmt.Printf("Failed to run server: %v\n", err)
    } else {
        fmt.Println("Server running on port 8080")
    }
}