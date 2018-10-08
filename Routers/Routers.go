package Routers

import (
	"github.com/gin-gonic/gin"
	"github.com/saseke/go-web/Controllers"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/item")
	{
		v1.GET("/list/page", Controllers.ListItemsPage)
		v1.GET("/list", Controllers.ListItems)
		v1.DELETE("/:id", Controllers.DeleteItem)
		v1.POST("", Controllers.AddNewItem)
		v1.PUT("", Controllers.PutItem)
	}
	r.GET("/test", Controllers.TestA)
	return r
}
