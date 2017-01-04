package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var WelcomeMessage string = "Hello"

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", handler2)
	http.ListenAndServe(":8080", nil)

}

// func handler(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprintf(w, "Hello World!\n")

// }

// func handler2(w http.ResponseWriter, r *http.Request) {

// 	//fmt.Fprintf(w, "Hello Earth\n")

// }

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type",
		"text/html")

	docBytes, err := ioutil.ReadFile(".\\add-content.html")
	if err != nil {
		panic(err)
	}

	docString := string(docBytes)

	tmpl, err := template.New("Add-Content").Parse(docString)
	if err == nil {
		tmpl.Execute(w, WelcomeMessage)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")

	docBytes, err := ioutil.ReadFile(".\\index.html")
	docString := string(docBytes)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("HomePage").Parse(docString)
	if err == nil {
		tmpl.Execute(w, WelcomeMessage)
	}

}
