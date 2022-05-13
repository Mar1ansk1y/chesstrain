package filevideo

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var ErrNoRecord = errors.New("video: подходящего видео не найдено")

type Video struct {
	gorm.Model
	ID       int
	Title    string
	Content  []byte
	Data     time.Time
	ModuleID uint
}
