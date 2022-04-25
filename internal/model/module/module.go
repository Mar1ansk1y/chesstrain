package module

import (
	"gorm.io/gorm"
	"time"
)

type Module struct {
	gorm.Model
	Files []byte
	Links string
	Text  string
	Date  time.Time
}
