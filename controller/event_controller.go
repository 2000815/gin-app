package controller

import (
	"my-go-app/model"
	"my-go-app/utility"

	"github.com/gin-gonic/gin"
)

func ShowAllBlog(c *gin.Context) {
	var events []model.Event
	result := utility.Db.Find(&events)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "データ取得エラー"})
		return
	}

	// データが空の場合の処理を追加する
	if len(events) == 0 {
		c.JSON(200, gin.H{"message": "データが見つかりません"})
		return
	}

	c.JSON(200, gin.H{"datas": events})
}