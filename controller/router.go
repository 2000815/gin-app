package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", ShowAllBlog)
	r.GET("/sss", ShowAllBlog)
	return r
}