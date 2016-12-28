package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var WelcomeMessage string = "Welcome to my website"

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/earth", handler2)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")

	docBytes, err := ioutil.ReadFile("C:\\GoRepo\\src\\github.com\\iggyuga\\httpserver2\\templates\\index.html")
	docString := string(docBytes)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("HomePage").Parse(docString)
	if err == nil {
		tmpl.Execute(w, WelcomeMessage)
	}

	//fmt.Fprintf(w, "Hello World!\n")

}

func handler2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Earth\n")

}
