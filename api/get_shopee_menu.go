package api

import (
	"encoding/json"
	m "github/free-order-be/api/middleware"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s Server) GetShopeeMenu(w http.ResponseWriter, r *http.Request) {
	res, err := s.logic.GetShopeeMenu(r.Context(), r.URL.Query().Get("endpoint"))
	if err != nil {
		logrus.Errorf("[API] GetShopeeMenu on err: %s \n", err)
		m.InternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(res)
}
