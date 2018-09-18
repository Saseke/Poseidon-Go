package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/saseke/go-web/Config"
)
import "time"

type Item struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	SellPoint string `json:"sell_point"`
	Price     float32
	Num       int
	LimitNum  int
	Image     string
	// 外键
	CId     int
	Status  int
	Created time.Time
	Updated time.Time

	SortOrder   int
	Remark      string
	Description string
}

func GetAllItems(items *[]Item) (err error) {
	if err = Config.DB.Find(items).Error; err != nil {
		return err
	}
	fmt.Print(items)
	return nil
}

func AddNewItem(item *Item) (err error) {
	if err := Config.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func GetOneItem(item *Item, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(item).Error; err != nil {
		return err
	}
	return nil
}

func PutOneItem(item *Item) (err error) {
	fmt.Println(item)
	Config.DB.Save(item)
	return nil
}

func DeleteItem(item *Item, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(item)
	return nil
}
