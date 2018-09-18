package ApiHelper

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Meta   interface{}
	Data   interface{}
}

func RespondJSON(c *gin.Context, status int, payload interface{}) {
	fmt.Println("status", status)
	var res ResponseData

	res.Status = status
	res.Data = payload
	c.JSON(200, res)
}
