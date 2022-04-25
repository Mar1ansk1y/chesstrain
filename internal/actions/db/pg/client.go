package pg

import (
	"ChessTrain/internal/model/bill"
	"ChessTrain/internal/model/course"
	"ChessTrain/internal/model/module"
	"ChessTrain/internal/model/user"
	"github.com/jinzhu/gorm"
	"log"
)

type SnippetModel struct {
	dbase *gorm.DB
}

func Init() *gorm.DB {
	connStr := "user=urfu password=123456 dbname=postgres sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&user.User{}, &bill.Bill{}, course.Course{}, module.Module{})
	return db
}

func (m *SnippetModel) Auth(nick, pass string) error {
	// Данные берутся из формы авторизации
	nick = "something"
	pass = "password"

	return nil
}
