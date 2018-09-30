package Models

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/saseke/go-web/Config"
	"math"
)
import "time"

type Item struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	SellPoint string  `json:"sell_point"`
	Price     float32 `json:"price"`
	Num       int     `json:"num"`
	LimitNum  int     `json:"limit_num"`
	Image     string  `json:"image"`
	// 外键
	CId     int       `json:"c_id"`
	Status  int       `json:"status"`
	Created time.Time `gorm:"default:'' "`
	Updated time.Time `gorm:"default:''"`

	SortOrder   int    `json:"sort_order"`
	Remark      string `json:"remark"`
	Description string `json:"description"`
}

func GetAllItems(items *[]Item) (err error) {
	if err = Config.DB.Find(items).Error; err != nil {
		return err
	}
	fmt.Print(items)
	return nil
}

// order : paging order, prop : paging property ,page : cur page ,limit : number of sheets displayed per page
func GetAllItemsPage(pageRows *PageRows, items *[]Item, orderBy string, page, limit int) (err error) {
	if err := Config.DB.Limit(limit).Offset((page - 1) * limit).Order(orderBy).Find(items).Error; err != nil {
		return err
	}
	PageRowsInit(pageRows, page+1, int(math.Ceil(float64(getItemsCount()/limit))), items)
	return nil
}

func getItemsCount() (count int) {
	Config.DB.Table("db_item").Count(&count)
	return
}

func AddNewItem(item *Item) (err error) {
	if err := Config.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func GetOneItem(item *Item, id string) (err error) {
	/*	if err := Config.DB.Where("id = ?", id).First(item).RecordNotFound(); err != nil {
			return err
		}
		return nil*/
	if Config.DB.Where("id = ?", id).First(item).RecordNotFound() {
		return errors.New("no this item")
	}
	return nil
}

func PutOneItem(item *Item) (err error) {
	if err := Config.DB.Model(item).Update(item).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItem(id string) (err error) {
	if err := Config.DB.Where("id = ?", id).Delete(&Item{}).Error; err != nil {
		return err;
	}
	return nil
}
