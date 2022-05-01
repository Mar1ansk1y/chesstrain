package handlers

import (
	video "ChessTrain/internal/actions/videos"
	"flag"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

type application struct {
	videos *video.Videomodel
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func Handle() {

	dsn := flag.String("dsn", "user=urfu password=123456 dbname=postgres sslmode=disable", "Название PSQL источника данных")

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := &application{
		videos: &video.Videomodel{DB: db},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/showcourse", showCourse)
	mux.HandleFunc("/profiles", profiles)
	mux.HandleFunc("/bill_page", bill_page)
	mux.HandleFunc("/hometask", app.hometask)
	mux.HandleFunc("/signup", signUp)
	mux.HandleFunc("/signin", signIn)
	mux.HandleFunc("/video", app.video)

	log.Println("запуск сервера на http://127.0.0.1:4000")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
