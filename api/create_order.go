package api

import (
	"encoding/json"
	"fmt"
	. "github/free-order-be/api/middleware"
	"github/free-order-be/models"

	"net/http"
)

func (s Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	req := &models.Order{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Printf("[API] CreateOrder on err: %s \n", err)
		BadRequest(w, err)
		return
	}

	resp, err := s.logic.CreateOrder(r.Context(), req)
	if err != nil {
		fmt.Printf("[API] CreateOrder on err: %s \n", err)
		InternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
