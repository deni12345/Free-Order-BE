package api

import (
	"fmt"
	. "github/lambda-microservice/api/middleware"
	"net/http"

	"github.com/gorilla/websocket"
)

func (s Server) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("[API] HandleWebSocket on err: %s", err)
		InternalError(w, err)
		return
	}
	s.connsMap[conn] = true
	go s.readLoop(conn)
}

func (s Server) readLoop(ws *websocket.Conn) {
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("[API] HandleWebSocket on err: %s", err)
			return
		}
		fmt.Printf("msg type is %v and msg is %v \n", msgType, msg)
		s.broadcast()
	}
}

func (s Server) broadcast() error {
	for conn := range s.connsMap {
		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello there"))
		if err != nil {
			return err
		}
	}
	return nil
}
