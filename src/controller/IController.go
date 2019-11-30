package controller

import "net/http"

type IController interface {
	Execute(w http.ResponseWriter, r *http.Request)
}
