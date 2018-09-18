package Models

import "github.com/saseke/go-web/Config"

type Article struct {
	ID      int
	Title   string
	Content string
}

func GetAllArticles(articles *[]Article) (err error) {
	if err = Config.DB.Find(articles).Error; err != nil {
		return err
	}
	return nil
}
