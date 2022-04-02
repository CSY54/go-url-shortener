// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/CSY54/go-url-shortener/src/url"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func initializeUrlController(db *gorm.DB) url.UrlController {
	urlRepository := url.ProvideUrlRepository(db)
	urlService := url.ProvideUrlService(urlRepository)
	urlController := url.ProvideUrlController(urlService)
	return urlController
}