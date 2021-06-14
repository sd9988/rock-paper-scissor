package controllers

import (
	"fmt"
	"log"
	"net/http"
	"rps/api/middlewares"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middlewares.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}
func (server *Server) initializeRoutes() {
	server.Router.Use(middlewares.CORS)
	server.setJSON("/health", server.Health, http.MethodGet)
	server.setJSON("/choices", server.Choices, http.MethodGet)
	server.setJSON("/choice", server.Choice, http.MethodGet)
	server.setJSON("/play", server.Play, http.MethodPost)
}

func (server *Server) Run(addr string) {
	fmt.Println("Running")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
