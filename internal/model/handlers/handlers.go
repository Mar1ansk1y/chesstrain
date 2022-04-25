package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{"./pkg/templates/home.html",
		"./pkg/templates/base.layot.html"} // html файлы домашней страницы
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

func showCourse(w http.ResponseWriter, r *http.Request) {
	files := []string{"./pkg/templates/courses.html",
		"./pkg/templates/base.layot.html"} // html файлы домашней страницы
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

func profiles(w http.ResponseWriter, r *http.Request) {

	files := []string{"./pkg/templates/NotAuthorized.html",
		"./pkg/templates/base.layot.html"} // html файлы домашней страницы
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

func bill_page(w http.ResponseWriter, r *http.Request) {
	files := []string{"./pkg/templates/bill_page.html",
		"./pkg/templates/base.layot.html"} // html файлы домашней страницы
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

func hometask(w http.ResponseWriter, r *http.Request) {
	files := []string{"./pkg/templates/hometask.html",
		"./pkg/templates/base.layot.html"} // html файлы домашней страницы
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

func signUp(w http.ResponseWriter, r *http.Request) {
	files := []string{"./pkg/templates/signup.html",
		"./pkg/templates/base.layot.html"} // html файлы домашней страницы
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

func signIn(w http.ResponseWriter, r *http.Request) {
	files := []string{"./pkg/templates/signin.html",
		"./pkg/templates/base.layot.html"} // html файлы домашней страницы
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
