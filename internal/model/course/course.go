package course

import "time"

type Course struct {
	Enable bool `json:"enable"`
	Links string `json:"links"`
	UserRefer uint
	DateStart time.Time `json:"date_start"`
	DateEnd time.Time `json:"date_end"`
}

func (c *Course) CourseEnable() {
//проверка токенов

	c.Enable = true
}