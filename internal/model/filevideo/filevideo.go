package filevideo

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("video: подходящего видео не найдено")

type Video struct {
	ID      int
	Title   string
	Content []byte
	Data    time.Time
}
