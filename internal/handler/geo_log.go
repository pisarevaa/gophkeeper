package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pisarevaa/fastlog/internal/model"
)

type QueryMetrics struct {
	ID    string `json:"id"`
	MType string `json:"type"`
}

// Отправка гео лога
// SendGeoLog godoc
//
//	@Summary	Regiser user
//	@Schemes
//	@Tags		Log
//	@Accept		json
//	@Produce	json
//	@Param		request	body		storage.RegisterUser	true	"Body"
//	@Success	200		{object}	storage.Success			"Response"
//	@Failure	409		{object}	storage.Error			"Login is already used"
//	@Failure	500		{object}	storage.Error			"Error"
//	@Router		/api/user/register [post]
func (s *Handler) SendGeoLog(w http.ResponseWriter, r *http.Request) {
	var logs model.GeoLog
	if err := json.NewDecoder(r.Body).Decode(&logs); err != nil {
		s.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if err := s.Validator.Struct(logs); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if _, err := w.Write([]byte("{\"data\":true}")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
