package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lowsound42/goweb/controllers"
	"github.com/lowsound42/goweb/models"
	"github.com/lowsound42/goweb/templates"
	"github.com/lowsound42/goweb/views"
)

func main() {
	r := chi.NewRouter()

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}

	// Setup our controllers
	usersC := controllers.Users{
		UserService: &userService,
	}

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.tmpl", "tailwind.tmpl"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.tmpl", "tailwind.tmpl"))))
	r.Get("/signup", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "signup.tmpl", "tailwind.tmpl")),
	))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.tmpl", "tailwind.tmpl"))))

	usersC.Templates.View = views.Must(views.ParseFS(templates.FS, "signup.tmpl", "tailwind.tmpl"))
	r.Get("/signup", usersC.View)
	r.Post("/signup", usersC.Create)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Ahoy matey, we be sailin' on port :3000")
	http.ListenAndServe(":3000", r)
}
