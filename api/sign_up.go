package api

import (
	"encoding/json"
	"fmt"
	. "github/free-order-be/api/middleware"
	"github/free-order-be/models"
	"net/http"
)

func (s Server) SignUp(w http.ResponseWriter, r *http.Request) {
	req := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Printf("[API] SignUp on err: %s \n", err)
		BadRequest(w, err)
		return
	}

	resp, err := s.logic.CreateUser(r.Context(), req)
	if err != nil {
		InternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
