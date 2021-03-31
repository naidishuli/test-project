package controllers

import (
	. "client_api/domain/interfaces"
	"github.com/gorilla/context"
	"net/http"
)

type Controller interface {
	Scope(*http.Request) AppScope
}


type RESTController interface {
	Controller
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Gets(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

func NewController() *controller {
	return &controller{}
}

func (c *controller) Scope(r *http.Request) AppScope {
	return context.Get(r, "appScope").(AppScope)
}
