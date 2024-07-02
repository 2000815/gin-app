package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type event struct {
    Id          int    `json:id`
    UserId     int    `json:user_id`
    Summary     string `json:summary`
    Dtstart     string `json:dtstart`
    Dtend       string `json:dtend`
    Description string `json:description`
    Year        int    `json:year`
    Month       int    `json:month`
    Day         int    `json:day`
  }


  func gormConnect() *gorm.DB {
    DBMS := "mysql"
    USER := "root"
    PASS := "password"
    HOST := "mysql"  // Docker ComposeのMySQLサービス名を指定する
    PORT := "3306"
    DBNAME := "mydatabase"
    CONNECT := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        fmt.Println("データベース接続エラー:", err.Error())
        panic(err.Error())
    }
    fmt.Println("データベースに接続しました:", db)
    return db
}

func main() {
    router := gin.Default()

    router.GET("/event/:id", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello world")

        fmt.Println("hello")
        db := gormConnect()


        // パラメータからイベントIDを取得
        id := c.Param("id")

        // データベースからイベントを取得
        eventEx := event{}
        idNum, _ := strconv.Atoi(id)
        eventEx.Id = idNum
        db.First(&eventEx)
        jsonData, err := json.Marshal(eventEx)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "JSON変換エラー"})
            return
        }

        // JSONデータをレスポンスとして返す
        c.Data(http.StatusOK, "application/json", jsonData)
    })

    // サーバーを起動
    router.Run(":8080")
}