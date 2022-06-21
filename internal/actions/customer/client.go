package customer

import (
	"ChessTrain/internal/model/bill"
	"ChessTrain/internal/model/course"
	"ChessTrain/internal/model/hometask"
	"ChessTrain/internal/model/module"
	"ChessTrain/internal/model/user"
	"errors"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SnippetModel struct {
	Dbase *gorm.DB
}

func Init() *gorm.DB {
	connStr := "host=80.87.192.62 port=32454 user=urfu password=123456 dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&user.User{}, &bill.Bill{}, course.Course{}, module.Module{})
	return db
}

func (s *SnippetModel) Insert(Text, Header string, dat time.Time) (int, error) {

	hometsk := hometask.Hometask{Header: Header, Text: Text, Date: dat}

	s.Dbase.Select("Title", "Content", "Data").Create(&hometsk)

	return hometsk.ID, nil
}

func (s *SnippetModel) Get(id int) (*hometask.Hometask, error) {

	var hometsk hometask.Hometask

	f := &hometask.Hometask{}

	row := s.Dbase.First(hometsk, id).Scan(&f)

	if row.Error != nil {
		if errors.Is(row.Error, gorm.ErrRecordNotFound) {
			return nil, hometask.ErrNoRecord
		} else {
			return nil, row.Error
		}
	}

	return f, nil
}

func (m *SnippetModel) Auth(nick, pass string) error {
	// Данные берутся из формы авторизации
	nick = "something"
	pass = "password"

	return nil
}
