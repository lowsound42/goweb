package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lowsound42/goweb/controllers"
	"github.com/lowsound42/goweb/templates"
	"github.com/lowsound42/goweb/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.tmpl", "tailwind.tmpl"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.tmpl", "tailwind.tmpl"))))
	r.Get("/signup", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "signup.tmpl", "tailwind.tmpl")),
	))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.tmpl", "tailwind.tmpl"))))

	var usersC controllers.Users
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.tmpl", "tailwind.tmpl"))
	r.Get("/signup", usersC.New)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Ahoy matey, we be sailin' on port :3000")
	http.ListenAndServe(":3000", r)
}
