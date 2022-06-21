package user

import (
	credit_card "ChessTrain/internal/model/credit-card"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname    string `json:"nickname"`
	Surname     string `json:"surname"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	CreditCards []credit_card.CreditCard
	SuperUser   bool `json:"superUser"`
	//CourseID    []course.Course
}
