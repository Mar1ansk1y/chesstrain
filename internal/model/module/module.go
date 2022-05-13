package module

import (
	"ChessTrain/internal/model/filevideo"
	"gorm.io/gorm"
	"time"
)

type Module struct {
	gorm.Model
	Files    []byte
	Links    string
	Text     string
	Date     time.Time
	Video    []filevideo.Video
	CourseID uint
}
