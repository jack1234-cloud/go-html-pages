package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
	case "/login-submit":
	default:
		fmt.Fprint(w, "Hi Jack")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}
