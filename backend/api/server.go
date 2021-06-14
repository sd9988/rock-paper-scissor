package api

import (
	"rps/api/controllers"

	"github.com/gorilla/mux"
)

func New() *controllers.Server {
	s := &controllers.Server{}
	s.Router = mux.NewRouter()
	return s
}

func Run() {
	s := New()
	s.Initialize()
	s.Run(":8080")
}
