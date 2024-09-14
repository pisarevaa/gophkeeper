package handler

import (
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

func (s *Handler) GetData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		utils.JSON(w, http.StatusInternalServerError, model.Error{Error: "Error to cast userID into int64"})
		return
	}
	data, status, err := s.KeeperService.GetData(r.Context(), userID)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	var ordersResponse []model.DataResponse
	for _, d := range data {
		ordersResponse = append(
			ordersResponse,
			model.DataResponse{
				ID:        d.ID,
				Data:      d.Data,
				Type:      d.Type,
				CreatedAt: model.DateTime(d.CreatedAt),
				UpdatedAt: model.DateTime(d.UpdatedAt),
			},
		)
	}
	utils.JSON(w, http.StatusOK, ordersResponse)
}

func (s *Handler) GetDataByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		utils.JSON(w, http.StatusInternalServerError, model.Error{Error: "Error to cast userID into int64"})
		return
	}
	dataID, err := utils.GetDataID(r)
	if err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	data, status, err := s.KeeperService.GetDataByID(r.Context(), userID, dataID)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	utils.JSON(w, http.StatusOK, model.DataResponse{
		ID:        data.ID,
		Data:      data.Data,
		Type:      data.Type,
		CreatedAt: model.DateTime(data.CreatedAt),
		UpdatedAt: model.DateTime(data.UpdatedAt),
	})
}
