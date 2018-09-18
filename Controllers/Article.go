package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saseke/go-web/ApiHelper"
	"github.com/saseke/go-web/Models"
)

func GetAllArticles(c *gin.Context) {
	var articles []Models.Article
	err := Models.GetAllArticles(&articles)
	if err != nil {
		ApiHelper.RespondJSON(c, 404, articles)
	} else {
		ApiHelper.RespondJSON(c, 200, articles)
	}
}
