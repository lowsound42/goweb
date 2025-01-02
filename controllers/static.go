package controllers

import (
	"html/template"
	"net/http"
)

func FAQ(tpl Executor) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Do you do business?",
			Answer:   "Yes! Business is done here.",
		},
		{
			Question: "What about customer service?",
			Answer:   "So much of it.",
		},
		{
			Question: "Give me your email address",
			Answer:   `okay fine, <a href="mailto:o.khandxb@gmail.com">o.khandxb@gmail.com</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}

func StaticHandler(tpl Executor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
