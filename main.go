package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
  <ul>
	<li>
	  <b>Is there a free version?</b>
	  Yes! We offer a free trial for 30 days on any paid plans.
	</li>
	<li>
	  <b>What are your support hours?</b>
	  We have support staff answering emails 24/7, though response
	  times may be a bit slower on weekends.
	</li>
	<li>
	  <b>How do I contact support?</b>
	  Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
	</li>
  </ul>
  `)
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("<h1>param finder %v </h1>", id)))
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:o.khandxb@gmail.com\">o.khandxb@gmail.com</a>.</p>")
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		panic(err) // TODO -> remove panic
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err) // TODO -> remove panic
	}

}

func main() {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/", homeHandler)
	})
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/galleries/{id}", paramHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Ahoy matey, we be sailin' on port :3000")
	http.ListenAndServe(":3000", r)
}
