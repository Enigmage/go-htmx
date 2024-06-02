package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
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
			{Title: "Harry Potter", Director: "Tyson Jackson"},
		},
	}
	tmpl.Execute(resp, films)
}

func main() {
	fmt.Println("Hello World")
	// var whoamiHandler http.HandlerFunc = func(resp http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(resp, req.Host)
	// }
	whoamiHandler := func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, req.Host)
	}
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/whoami", whoamiHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
