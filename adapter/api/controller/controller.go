package controller

import "net/http"

type Controller interface {
	Execute(w http.ResponseWriter, r *http.Request)
}