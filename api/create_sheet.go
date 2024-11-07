package api

import (
	"encoding/json"
	"fmt"
	. "github/free-order-be/api/middleware"
	"github/free-order-be/models"

	"net/http"
)

func (s Server) CreateSheet(w http.ResponseWriter, r *http.Request) {
	req := &models.Sheet{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Printf("[API] CreateSheet on err: %s \n", err)
		BadRequest(w, err)
		return
	}

	resp, err := s.logic.CreateSheet(r.Context(), req)
	if err != nil {
		fmt.Printf("[API] CreateSheet on err: %s \n", err)
		InternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
