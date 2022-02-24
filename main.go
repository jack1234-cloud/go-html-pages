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
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
}

var userDB = map[string]string{
	"jack": "password",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User logged in successfully")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found in DB")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	default:
		fmt.Fprintf(w, "Hi Jack")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
	// http.ListenAndServeTLS("", "cert.pem", "key.pem", nil)
}
