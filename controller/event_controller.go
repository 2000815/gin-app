package controller

import (
	"my-go-app/model"
	"my-go-app/utility"
	"my-go-app/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var events []model.Event
	result := utility.Db.Find(&events)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "データ取得エラー"})
		return
	}

	if len(events) == 0 {
		c.JSON(200, gin.H{"message": "データが見つかりません"})
		return
	}

	jsonData := view.ProcessEvents(events)
	c.JSON(http.StatusOK, jsonData)
}