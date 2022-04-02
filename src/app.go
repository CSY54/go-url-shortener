package app

import (
	"github.com/CSY54/go-url-shortener/src/url"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupDatabase(db *gorm.DB) {
	db.AutoMigrate(&url.Url{})
}

func teardownDatabase(db *gorm.DB) {
	if err := db.Migrator().DropTable(&url.Url{}); err != nil {
		panic(err)
	}
}

func setupRoute(r *gin.Engine, db *gorm.DB) {
	urlController := initializeUrlController(db)

	api := r.Group("/api")
	{
		apiV1 := api.Group("/v1")
		{
			apiV1.POST("/urls", urlController.UploadUrl)
		}
	}

	r.GET("/:urlId", urlController.RedirectUrl)
}

func Init() *gin.Engine {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	setupDatabase(db)

	r := gin.Default()

	setupRoute(r, db)

	return r
}

func Run(r *gin.Engine) {
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
