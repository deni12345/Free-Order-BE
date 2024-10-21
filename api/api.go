package api

import (
	. "github/free-order-be/api/middleware"
	"github/free-order-be/internal/logic"

	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"golang.org/x/oauth2"
)

type Server struct {
	Router      *mux.Router
	logic       logic.Logic
	googleOauth *oauth2.Config
	upgrader    websocket.Upgrader
	connsMap    map[*websocket.Conn]bool
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		//logic:    logic.NewLogicImpl(),
		upgrader: configWebSocketUpgrader(),
	}

	s.Router.Use(ContentTypeJsonMiddleware)
	mount(s.Router, "/api/public", s.publicAPI())
	mount(s.Router, "/api/private", s.privateAPI())
	mount(s.Router, "/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not Found"))
	}))
	return s
}

func (s *Server) publicAPI() http.Handler {
	router := mux.NewRouter()
	router.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}).Methods("GET")

	// //Sign in/up with official account
	// router.HandleFunc("/sign-in", s.SignIn).Methods("POST")
	// router.HandleFunc("/sign-up", s.SignUp).Methods("POST")

	// //Sign in/up with google account
	// router.HandleFunc("/auth/google/login", s.GoogleSignIn).Methods("GET")
	// router.HandleFunc("/auth/google/callback", s.GoogleCallBack).Methods("GET")

	// //Sheet routes
	// router.HandleFunc("/sheet", s.CreateSheet).Methods("POST")

	// router.HandleFunc("/ws", s.HandleWebSocket).Methods("GET")
	return router
}

func (s *Server) privateAPI() http.Handler {
	router := mux.NewRouter()
	return router
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}

func configWebSocketUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
