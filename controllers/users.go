package controllers

import (
	"fmt"
	"net/http"

	"github.com/lowsound42/goweb/models"
)

type Users struct {
	Templates struct {
		View Executor
	}
	UserService *models.UserService
}

func (u *Users) View(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.View.Execute(w, data)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %+v", user)
}
