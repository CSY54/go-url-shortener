package url

import (
	"gorm.io/gorm"
)

type UrlRepository struct {
	DB *gorm.DB
}

func ProvideUrlRepository(DB *gorm.DB) UrlRepository {
	return UrlRepository{DB: DB}
}

func (u *UrlRepository) Create(url Url) Url {
	u.DB.Create(&url)

	return url
}

func (u *UrlRepository) FindByID(id int) Url {
	var url Url
	u.DB.Limit(1).Find(&url, id)

	return url
}
