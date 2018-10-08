package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saseke/go-web/ApiHelper"
	"net/http"
)

func TestA(c *gin.Context){
	ApiHelper.RespondJSON(c,http.StatusOK,"test")
}
