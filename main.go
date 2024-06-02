package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func indexHandler(resp http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	// syntax to declare and initialize map in same line
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
		},
	}
	tmpl.Execute(resp, films)
}

func addFilms(resp http.ResponseWriter, req *http.Request) {
    time.Sleep(2 * time.Second)
    log.Print("htmx requerst received")
    title := req.PostFormValue("title")
    director := req.PostFormValue("director")

    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.ExecuteTemplate(resp, "film-list-element", Film{Title: title, Director: director})
}

func main() {
    fmt.Println("Starting server at port 8000")
	// var whoamiHandler http.HandlerFunc = func(resp http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(resp, req.Host)
	// }
	whoamiHandler := func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, req.Host)
	}
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/whoami", whoamiHandler)
	http.HandleFunc("/add-film", addFilms)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
