package video

import (
	"ChessTrain/internal/model/filevideo"
	"errors"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Videomodel struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	connStr := "host=80.87.192.62 port=32454 user=urfu password=123456 dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(filevideo.Video{})
	return db
}

func (v *Videomodel) Insert(titl string, dat time.Time, cont []byte) (int, error) {

	videos := filevideo.Video{Title: titl, Content: cont, Data: dat}

	v.DB.Select("Title", "Content", "Data").Create(&videos)

	return videos.ID, nil
}

func (v *Videomodel) Get(id int) (*filevideo.Video, error) {

	var videos filevideo.Video

	f := &filevideo.Video{}

	row := v.DB.First(&videos, id).Scan(&f)

	if row.Error != nil {
		if errors.Is(row.Error, gorm.ErrRecordNotFound) {
			return nil, filevideo.ErrNoRecord
		} else {
			return nil, row.Error
		}
	}

	return f, nil
}
