package api

import (
	. "github/lambda-microservice/api/middleware"
	"github/lambda-microservice/internal/logic"
	"github/lambda-microservice/internal/oauth2"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
	mount(r, "/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("default route")
	}))
	s.Mux = r

	return s
}

func (s *Server) publicAPI() http.Handler {
	router := mux.NewRouter()
	router.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	router.Path("/sign-in").HandlerFunc(s.SignIn).Methods("POST")
	router.HandleFunc("/sign-up", s.SignUp).Methods("POST")
	router.HandleFunc("/auth/google/login", oauth2.OauthGoogleLogin).Methods("GET")
	router.HandleFunc("/auth/google/callback", oauth2.OauthGoogleCallBack).Methods("GET")
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
