package api

import (
	"encoding/json"
	m "github/free-order-be/api/middleware"
	"net/http"
)

func (s Server) GetShopeeMenu(w http.ResponseWriter, r *http.Request) {
	res, err := s.logic.GetShopeeMenu(r.Context(), r.URL.Query().Get("endpoint"))
	if err != nil {
		m.InternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(res)
}
