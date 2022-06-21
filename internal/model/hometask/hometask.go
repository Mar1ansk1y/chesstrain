package hometask

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var ErrNoRecord = errors.New("hometask: Данного задания не существует")

type Hometask struct {
	gorm.Model
	ID     int
	Header string
	Text   string
	Date   time.Time
}
