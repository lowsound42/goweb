package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/lowsound42/goweb/controllers"
	"github.com/lowsound42/goweb/migrations"
	"github.com/lowsound42/goweb/models"
	"github.com/lowsound42/goweb/templates"
	"github.com/lowsound42/goweb/views"
)

func main() {
	r := chi.NewRouter()
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false),
	)

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}
	// Setup our controllers
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.tmpl", "tailwind.tmpl"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.tmpl", "tailwind.tmpl"))))
	r.Get("/signup", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "signup.tmpl", "tailwind.tmpl")),
	))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.tmpl", "tailwind.tmpl"))))

	usersC.Templates.SignUp = views.Must(views.ParseFS(templates.FS, "signup.tmpl", "tailwind.tmpl"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.tmpl", "tailwind.tmpl"))
	r.Get("/signin", usersC.SignIn)
	r.Get("/signup", usersC.SignUp)
	r.Post("/signup", usersC.Create)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)
	r.Post("/signout", usersC.ProcessSignOut)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Ahoy matey, we be sailin' on port :3000")
	http.ListenAndServe(":3000", csrfMw(r))
}
