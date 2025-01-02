package controllers

import "net/http"

type Executor interface {
	Execute(w http.ResponseWriter, data interface{})
}
