package user

import (
	"ChessTrain/internal/model/course"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname  string          `json:"nickname"`
	Surname   string          `json:"surname"`
	Name      string          `json:"name"`
	Password  string          `json:"password"`
	Courses   []course.Course `json:"course_._course" gorm:"foreignKey:userRefer"`
	SuperUser bool
}
