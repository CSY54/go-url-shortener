package url

import (
	"time"
)

type Url struct {
	ID int `gorm:"primaryKey"`
	Url string
	ExpireAt time.Time
}
