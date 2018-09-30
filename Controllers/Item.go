package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saseke/go-web/ApiHelper"
	"github.com/saseke/go-web/Models"
	"net/http"
	"strconv"
)

func ListItems(c *gin.Context) {
	var items []Models.Item
	err := Models.GetAllItems(&items)
	fmt.Println(items)
	if err != nil {
		ApiHelper.RespondJSON(c, http.StatusBadRequest, items)
	} else {
		ApiHelper.RespondJSON(c, http.StatusOK, items)
	}
}

// format : url?prop=**&order=(ASC/DESC)&page=*&limit=*
func ListItemsPage(c *gin.Context) {
	var items []Models.Item
	var pageRows Models.PageRows
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	prop := c.Query("prop")
	order := c.Query("order")
	orderBy := prop + " " + order
	err := Models.GetAllItemsPage(&pageRows, &items, orderBy, page, limit)
	if err != nil {
		ApiHelper.RespondJSON(c, 404, nil)
	} else {
		ApiHelper.RespondJSON(c, 200, pageRows)
	}
}
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	if err := Models.DeleteItem(id); err != nil {
		ApiHelper.RespondJSON(c, 404, nil)
	} else {
		ApiHelper.RespondJSON(c, 200, "Success")
	}
}

func GetOneItem(c *gin.Context) {
	var item Models.Item
	id := c.Param("id")
	if err := Models.GetOneItem(&item, id); err != nil {
		ApiHelper.RespondJSON(c, 404, nil)
	} else {
		ApiHelper.RespondJSON(c, 200, item)
	}
}

func AddNewItem(c *gin.Context) {
	var item Models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		ApiHelper.RespondJSON(c, http.StatusBadRequest, nil)
	} else {
		if err := Models.AddNewItem(&item); err != nil {
			ApiHelper.RespondJSON(c, http.StatusBadRequest, nil)
		} else {
			ApiHelper.RespondJSON(c, http.StatusOK, item)
		}
	}
}

func PutItem(c *gin.Context) {
	var item Models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		ApiHelper.RespondJSON(c, http.StatusBadRequest, nil)
	} else {
		if err := Models.PutOneItem(&item); err != nil {
			ApiHelper.RespondJSON(c, http.StatusBadRequest, nil)
		} else {
			ApiHelper.RespondJSON(c, http.StatusOK, item)
		}
	}
}
