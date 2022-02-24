package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error occurred parsing the file", err)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
	case "/login-submit":
	default:
		fmt.Fprint(w, "Hi Jack")
	}
}

func main() {
	http.HandleFunc("/", login)
	http.ListenAndServe("", nil)
}
