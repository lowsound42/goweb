package main

import (
	"fmt"
	"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:o.khandxb@gmail.com\">o.khandxb@gmail.com</a>.</p>")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my space</h1>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/path", pathHandler)
	fmt.Println("Ahoy matey, we be sailin' on port :3000")
	http.ListenAndServe(":3000", nil)
}
