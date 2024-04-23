package api

import (
	"encoding/json"
	"github/lambda-microservice/model"
	"log"
	"net/http"
)

func (s Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Fatalf("API CreateUser on err: %s", err)
		BadRequest(w, err)
	}
	resp, err := s.logic.CreateUser(req)
	if err != nil {
		log.Fatalf("API CreateUser on err: %s", err)
		InternalError(w, err)
	}
	json.NewEncoder(w).Encode(resp)
}
