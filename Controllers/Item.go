package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saseke/go-web/ApiHelper"
	"github.com/saseke/go-web/Models"
)

func ListItems(c *gin.Context) {
	var items []Models.Item
	err := Models.GetAllItems(&items)
	fmt.Println(items)
	if err != nil {
		ApiHelper.RespondJSON(c, 404, items)
	} else {
		ApiHelper.RespondJSON(c, 200, items)
	}
}
