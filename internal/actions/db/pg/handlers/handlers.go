package handlers

import (
	"ChessTrain/internal/model/filevideo"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
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

func showCourse(w http.ResponseWriter, r *http.Request) {
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

func profiles(w http.ResponseWriter, r *http.Request) {

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

func bill_page(w http.ResponseWriter, r *http.Request) {
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

func (app *application) hometask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("MyFile")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error Retrieving the File", http.StatusNotImplemented)
	}

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error!", http.StatusBadGateway)
	}

	id, err := app.videos.Insert(handler.Filename, time.Now().UTC(), fileBytes)

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

	http.Redirect(w, r, fmt.Sprintf("/video?id=%d", id), http.StatusSeeOther)
}

func signUp(w http.ResponseWriter, r *http.Request) {
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

func (app *application) video(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	v, err := app.videos.Get(id)
	if err != nil {
		if errors.Is(err, filevideo.ErrNoRecord) {
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

func signIn(w http.ResponseWriter, r *http.Request) {
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
