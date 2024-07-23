package api

import (
	"github/lambda-microservice/api/middleware"
	"github/lambda-microservice/internal/logic"

	"github.com/gorilla/mux"
)

type Server struct {
	Mux   *mux.Router
	logic logic.Logic
}

func NewServer() *Server {
	l := logic.NewLogicImpl()
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(middleware.ContentTypeJsonMiddleware)

	s := &Server{
		logic: l,
		Mux:   r,
	}
	s.publicAPI()
	return s
}

func (s *Server) publicAPI() {
	sub := s.Mux.PathPrefix("/public").Subrouter()
	sub.HandleFunc("/sign-in", s.SignIn).Methods("POST")
	sub.HandleFunc("/sign-up", s.SignUp).Methods("POST")
}

// func (s *Server) privateAPI() http.Handler {
// 	h := mux.NewRouter()
// 	h.HandleFunc("/sign-in", s.SignUp).Methods("GET")
// 	return h
// }
