//+build wireinject

package app

import (
	"github.com/CSY54/go-url-shortener/src/url"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func initializeUrlController(db *gorm.DB) url.UrlController {
	wire.Build(
		url.ProvideUrlRepository,
		url.ProvideUrlService,
		url.ProvideUrlController,
	)

	return url.UrlController{}
}
