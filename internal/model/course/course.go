package course

import (
	"ChessTrain/internal/model/module"
	"ChessTrain/internal/model/user"

	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Enable    bool            `json:"enable"`
	Links     string          `json:"links"`
	Users     []user.User     `json:"users"`
	Modules   []module.Module `json:"modules"`
	DateStart time.Time       `json:"date_start"`
	DateEnd   time.Time       `json:"date_end"`
}

func (c *Course) CourseEnable() {
	//проверка токенов

	c.Enable = true
}
