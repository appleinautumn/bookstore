package server

import (
	"fmt"
	"net/http"

	"gotu/bookstore/internal/handler"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(apiPublicHandler *handler.ApiHandler) *Server {
	router := NewRouter(apiPublicHandler)
	srv := &Server{
		router: router,
	}
	return srv
}

func (s *Server) Start(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.router)
}
