package api

import (
	"encoding/json"
	"fmt"
	. "github/lambda-microservice/api/middleware"
	"github/lambda-microservice/models"
	"net/http"
)

func (s Server) SignIn(w http.ResponseWriter, r *http.Request) {
	req := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Printf("[API] SignIn on err: %s \n", err)
		BadRequest(w, err)
		return
	}

	resp, err := s.logic.SignIn(r.Context(), req)
	if err != nil {
		fmt.Printf("[API] SignIn on err: %s \n", err)
		InternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
