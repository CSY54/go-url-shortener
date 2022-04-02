package url

import "time"

type UploadUrlDTO struct {
	Url string `json:"url" binding:"required"`
	ExpireAt time.Time `json:"expireAt" binding:"required"`
}

type ResponseUrlDTO struct {
	ID string `json:"id"`
	ShortUrl string `json:"shortUrl"`
}

type RedirectUrlDTO struct {
	ID string `uri:"urlId" binding:"required"`
}
