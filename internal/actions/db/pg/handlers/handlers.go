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

	files := []string{"./ui/templates/home.html",
		"./ui/templates/base_header.html"}
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
	files := []string{"./ui/templates/courses.html",
		"./ui/templates/base.layot.html"}
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

	files := []string{"./ui/templates/Authorized.html",
		"./ui/templates/base.layot.html"}
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
	files := []string{"./ui/templates/bill_page.html",
		"./ui/templates/base.layot.html"}
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

	files := []string{"./ui/templates/createhometask.html",
		"./ui/templates/base.layot.html"}

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

func (app *application) signUp(w http.ResponseWriter, r *http.Request) {
	files := []string{"./ui/templates/signup.html",
		"./ui/templates/base_header.html"}
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

func (app *application) save_hometask(w http.ResponseWriter, r *http.Request) {

	header := r.FormValue("header")
	text_hometask := r.FormValue("text_hometask")
	date := time.Now().UTC()

	id, err := app.client.Insert(text_hometask, header, date)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
	}

	http.Redirect(w, r, fmt.Sprintf("/hometask?id=%d", id), http.StatusSeeOther)
}

func (app *application) signIn(w http.ResponseWriter, r *http.Request) {
	files := []string{"./ui/templates/signin.html",
		"./ui/templates/base_header.html"}
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

func (app *application) AboutUs(w http.ResponseWriter, r *http.Request) {

	files := []string{"./ui/templates/AboutUs.html",
		"./ui/templates/base_header.html"}

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

func (app *application) AfterLogin(w http.ResponseWriter, r *http.Request) {

	files := []string{"./ui/templates/AfterLogin.html",
		"./ui/templates/base_header.html"}

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

	files := []string{"./ui/templates/show_hometask.html",
		"./ui/templates/base.layout.html"}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

	err = ts.Execute(w, v)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

}

func (app *application) Auth(w http.ResponseWriter, r *http.Request) {

	uname := r.FormValue("uname")
	psw := r.FormValue("psw")
	if uname == "admin" && psw == "password" {
		http.Redirect(w, r, "/afterlogin", http.StatusSeeOther)
	} else {
		http.Error(w, "Неверный логин или пароль", 500)
	}
}
