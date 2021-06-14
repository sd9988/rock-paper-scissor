package controllers

import (
	"github.com/gorilla/mux"
)

func (server *Server) Initialize() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}
