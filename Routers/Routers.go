package Routers

import (
	"github.com/gin-gonic/gin"
	"github.com/saseke/go-web/Controllers"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/item")
	{
		v1.GET("/list", Controllers.ListItems)
	}
	return r
}
