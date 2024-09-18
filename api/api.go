package api

import (
	. "github/lambda-microservice/api/middleware"
	"github/lambda-microservice/internal/logic"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	Router   *mux.Router
	logic    logic.Logic
	upgrader websocket.Upgrader
	connsMap map[*websocket.Conn]bool
}

func NewServer() *Server {
	s := &Server{
		Router:   mux.NewRouter(),
		logic:    logic.NewLogicImpl(),
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
	router.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/sign-in", s.SignIn).Methods("POST")
	router.HandleFunc("/sign-up", s.SignUp).Methods("POST")
	router.HandleFunc("/auth/google/login", s.OauthGoogleLogin).Methods("GET")
	router.HandleFunc("/auth/google/callback", s.OauthGoogleCallBack).Methods("GET")
	router.HandleFunc("/ws", s.HandleWebSocket).Methods("GET")
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
