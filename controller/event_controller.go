package controller

import (
	"github.com/gin-gonic/gin"
)

func ShowAllBlog(c *gin.Context) {
	// datas := model.GetAll()
	// jsonData, err := json.Marshal(datas)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "JSON変換エラー"})
	// 	return
	// }

  //       // JSONデータをレスポンスとして返す
	// c.Data(http.StatusOK, "application/json", jsonData)
	// // c.HTML(200, "index.html", gin.H{"datas": datas})

	c.JSON(200, gin.H{"message": "リダイレクトしました"})
}