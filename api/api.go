package api

import (
	. "github/free-order-be/api/middleware"
	"github/free-order-be/config"
	"github/free-order-be/internal/dao"
	"github/free-order-be/internal/logic"

	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	people "google.golang.org/api/people/v1"
)

//	@title			Free Order APIs
//	@version		1.0
//	@description	```{
// Free Order APIs containt api using for web application with the same name as Free Order,
// which is using for placing drinks and foods }``

// @contact.name				LiemDeni
// @contact.email				liemtran1414@gmail.com
// @host:						localhost:8080
// @BasePath:					/api
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
type Server struct {
	Router      *mux.Router
	logic       logic.Logic
	googleOauth *oauth2.Config
	upgrader    websocket.Upgrader
	connsMap    map[*websocket.Conn]bool
}

func NewServer(daoInst *dao.DAO) *Server {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/public/auth/google/callback",
		ClientID:     config.Values.GoogleID,
		ClientSecret: config.Values.GoogleClientSecret,
		Scopes:       []string{people.UserinfoEmailScope, people.UserinfoProfileScope},
		Endpoint:     google.Endpoint,
	}

	s := &Server{
		Router:      mux.NewRouter(),
		logic:       logic.NewLogicImpl(daoInst),
		upgrader:    configWebSocketUpgrader(),
		googleOauth: googleOauthConfig,
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
	router.HandleFunc("/users", s.GetUser).Methods("GET")

	// //Sign in/up with official account
	router.HandleFunc("/sign_in", s.SignIn).Methods("POST")
	router.HandleFunc("/sign_up", s.SignUp).Methods("POST")

	// //Sign in/up with google account
	router.HandleFunc("/auth/google/login", s.GoogleSignIn).Methods("GET")
	router.HandleFunc("/auth/google/callback", s.GoogleCallBack).Methods("GET")

	//Sheet routes
	router.HandleFunc("/sheets", s.CreateSheet).Methods("POST")
	router.HandleFunc("/sheets", s.GetSheet).Methods("GET")

	//Order routes
	router.HandleFunc("/orders", s.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", s.GetSheetOrders).Methods("GET")
	router.HandleFunc("/orders/user", s.GetUserOrders).Methods("GET")

	//Shopee routes
	router.HandleFunc("/shopee_menu", s.GetShopeeMenu).Methods("GET")

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
