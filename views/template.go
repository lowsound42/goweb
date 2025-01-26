package views

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/lowsound42/goweb/context"
	"github.com/lowsound42/goweb/models"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

type public interface {
	Public() string
}

type Template struct {
	htmlTpl *template.Template
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	htmlTpl := template.New(patterns[0])
	htmlTpl = htmlTpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField not implemented")
			},
			"currentUser": func() (*models.User, error) {
				return nil, fmt.Errorf("currentUser not implemented")
			},
			"errors": func() []string {
				return nil
			},
		},
	)
	htmlTpl, err := htmlTpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTpl: htmlTpl,
	}, nil
}

func errMessages(errs ...error) []string {
	var msgs []string
	for _, err := range errs {
		var pubErr public
		if errors.As(err, &pubErr) {
			msgs = append(msgs, pubErr.Public())
		} else {
			fmt.Println(err)
			msgs = append(msgs, "Something went wrong.")
		}
	}
	return msgs
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error) {
	htmlTpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}

	errMsgs := errMessages(errs...)
	htmlTpl = htmlTpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"currentUser": func() *models.User {
				return context.User(r.Context())
			},
			"errors": func() []string {
				return errMsgs
			},
		},
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var buf bytes.Buffer
	err = htmlTpl.Execute(&buf, data)

	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
