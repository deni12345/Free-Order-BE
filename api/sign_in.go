package api

import (
	"encoding/json"
	"github/lambda-microservice/models"
	"log"
	"net/http"
)

func (s Server) SignIn(w http.ResponseWriter, r *http.Request) {
	req := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("API SignIn on err: %s", err)
		BadRequest(w, err)
		return
	}

	resp, err := s.logic.SignIn(r.Context(), req)
	if err != nil {
		log.Printf("API SignIn on err: %s", err)
		InternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
