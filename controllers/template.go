package controllers

import "net/http"

type Executor interface {
	Execute(w http.ResponseWriter, r *http.Request, data interface{})
}
