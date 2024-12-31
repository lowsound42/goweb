package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my space</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Ahoy matey, we be sailin' on port :3000")
	http.ListenAndServe(":3000", nil)
}
