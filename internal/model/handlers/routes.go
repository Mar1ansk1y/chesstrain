package handlers

import (
	"log"
	"net/http"
)

func Handle() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/showcourse", showCourse)
	mux.HandleFunc("/profiles", profiles)
	mux.HandleFunc("/bill_page", bill_page)
	mux.HandleFunc("/hometask", hometask)
	mux.HandleFunc("/signup", signUp)
	mux.HandleFunc("/signin", signIn)

	log.Println("запуск сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
