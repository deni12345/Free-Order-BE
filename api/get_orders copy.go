package api

import (
	"encoding/json"
	. "github/free-order-be/api/middleware"
	"github/free-order-be/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s Server) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	req := &models.GetUserOrdersReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Infof("[API] GetUserOrders on err: %s \n", err)
		BadRequest(w, err)
		return
	}

	resp, err := s.logic.GetUserOrders(r.Context(), req)
	if err != nil {
		logrus.Infof("[API] GetUserOrders on err: %s \n", err)
		InternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
