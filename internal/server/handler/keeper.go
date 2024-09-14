package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"strconv"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

func (s *Handler) GetData(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	data, status, err := s.KeeperService.GetData(r.Context(), userID)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	utils.JSON(w, http.StatusOK, model.DataResponse{
		ID:        data.ID,
		Data:     data.Data,
		Type:     data.Type,
		CreatedAt: utils.Datetime(data.CreatedAt),
		UpdatedAt: utils.Datetime(data.UpdatedAt),
	})
}

func (s *Handler) GetDataByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	dataIDString := chi.URLParam(r, "dataID")
	if dataIDString == "" {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: "Path param dataID is not set"})
		return
	}
	dataID, err := strconv.ParseInt(dataIDString, 10, 64)
    if err != nil {
        utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: "Path param dataID is not integer"})
		return
    }
	data, status, err := s.KeeperService.GetDataByID(r.Context(), userID, dataID)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	utils.JSON(w, http.StatusOK, model.DataResponse{
		ID:        data.ID,
		Data:     data.Data,
		Type:     data.Type,
		CreatedAt: model.DateTime(data.CreatedAt),
		UpdatedAt: model.DateTime(data.UpdatedAt),
	})
}
