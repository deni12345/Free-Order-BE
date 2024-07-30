package api

import (
	. "github/lambda-microservice/api/middleware"
	"github/lambda-microservice/internal/logic"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Server struct {
	Mux   *mux.Router
	logic logic.Logic
}

func NewServer() *Server {
	l := logic.NewLogicImpl()
	s := &Server{
		logic: l,
	}

	r := mux.NewRouter()
	r.Use(ContentTypeJsonMiddleware)
	mount(r, "/api/public", s.publicAPI())
	mount(r, "/api/private", s.privateAPI())
	s.Mux = r

	return s
}

func (s *Server) publicAPI() http.Handler {
	router := mux.NewRouter()
	router.Path("/sign-in").HandlerFunc(s.SignIn).Methods("POST")
	router.HandleFunc("/sign-up", s.SignUp).Methods("POST")
	return router
}

func (s *Server) privateAPI() http.Handler {
	h := mux.NewRouter()
	return h
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
