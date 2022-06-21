package handlers

import (
	"ChessTrain/internal/model/hometask"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{"./pkg/templates/home.html",
		"./pkg/templates/base.layot.html"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}
}

func (app *application) showCourse(w http.ResponseWriter, r *http.Request) {
	files := []string{"./pkg/templates/courses.html",
		"./pkg/templates/base.layot.html"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}
}

func (app *application) profiles(w http.ResponseWriter, r *http.Request) {

	files := []string{"./pkg/templates/NotAuthorized.html",
		"./pkg/templates/base.layot.html"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

}

func (app *application) bill_page(w http.ResponseWriter, r *http.Request) {
	files := []string{"./pkg/templates/bill_page.html",
		"./pkg/templates/base.layot.html"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}
}

func (app *application) createhometask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	text := "Ваше домашнее задание"
	header := "Домашнее задание №1"
	date := time.Now().UTC()

	id, err := app.client.Insert(text, header, date)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
	}

	files := []string{"./pkg/templates/hometask.html",
		"./pkg/templates/base.layot.html"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

	http.Redirect(w, r, fmt.Sprintf("/hometask?id=%d", id), http.StatusSeeOther)
}

func (app *application) signUp(w http.ResponseWriter, r *http.Request) {
	files := []string{"./pkg/templates/signup.html",
		"./pkg/templates/base.layot.html"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}
}

func (app *application) hometask(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	v, err := app.client.Get(id)
	if err != nil {
		if errors.Is(err, hometask.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			log.Println(err.Error())
			http.Error(w, "error", 500)
		}
		return
	}

	files := []string{
		"./pkg/templates/video.html",
		"./pkg/templates/base.layot.html"}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		http.Error(w, "error", 500)
		return
	}

	err = ts.Execute(w, v)
	if err != nil {
		http.Error(w, "error", 500)
	}

}

func (app *application) signIn(w http.ResponseWriter, r *http.Request) {
	files := []string{"./pkg/templates/signin.html",
		"./pkg/templates/base.layot.html"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}
}
