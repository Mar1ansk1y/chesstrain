package handlers

import (
	"ChessTrain/internal/actions/customer"
	"flag"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type application struct {
	infolog  *log.Logger
	errorlog *log.Logger
	client   *customer.SnippetModel
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func Handle() {

	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	dsn := flag.String("dsn", "host=80.87.192.62 port=32454 user=urfu password=123456 dbname=postgres sslmode=disable", "Название PSQL источника данных")

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infolog:  infolog,
		errorlog: errorlog,
		client:   &customer.SnippetModel{Dbase: db},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/showcourse", app.showCourse)
	mux.HandleFunc("/profiles", app.profiles)
	mux.HandleFunc("/bill_page", app.bill_page)
	mux.HandleFunc("/hometask/create", app.createhometask)
	mux.HandleFunc("/signup", app.signUp)
	mux.HandleFunc("/signin", app.signIn)
	mux.HandleFunc("/hometask", app.hometask)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  mux,
	}

	infolog.Printf("Запуск сервера на %s", *addr)
	err = srv.ListenAndServe()
	errorlog.Fatal(err)
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database, err := db.DB()
	if err != nil {
		if err = database.Ping(); err != nil {
			return nil, err
		}
	}

	return db, nil
}
