package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"os"
)

func main() {

	log.Println("starting server")

	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./fonts"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	fs := http.FileServer(http.Dir("."))
	http.Handle("/style.css", fs)

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		tmpl.Execute(w, nil)
	})

	host := "127.0.0.1"
	//port := os.Getenv("PORT")
    port := "8080"

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
