package handler

import (
	"encoding/json"
	"io"
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

func (s *Handler) AddTextData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		utils.JSON(w, http.StatusInternalServerError, model.Error{Error: "Error to cast userID into int64"})
		return
	}
	var textData model.AddTextData
	if err := json.NewDecoder(r.Body).Decode(&textData); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	if err := s.Validator.Struct(textData); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	data, status, err := s.KeeperService.AddTextData(r.Context(), textData.Data, userID)
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

func (s *Handler) AddBinaryData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		utils.JSON(w, http.StatusInternalServerError, model.Error{Error: "Error to cast userID into int64"})
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, model.Error{Error: err.Error()})
		return
	}
	data, status, err := s.KeeperService.AddBinaryData(r.Context(), body, userID)
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

func (s *Handler) UpdateTextData(w http.ResponseWriter, r *http.Request) {
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
	var textData model.AddTextData
	if err := json.NewDecoder(r.Body).Decode(&textData); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	if err := s.Validator.Struct(textData); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	data, status, err := s.KeeperService.UpdateTextData(r.Context(), textData.Data, userID, dataID)
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

func (s *Handler) UpdateBinaryData(w http.ResponseWriter, r *http.Request) {
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, model.Error{Error: err.Error()})
		return
	}
	data, status, err := s.KeeperService.UpdateBinaryData(r.Context(), body, userID, dataID)
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

func (s *Handler) DeleteData(w http.ResponseWriter, r *http.Request) {
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
	data, status, err := s.KeeperService.DeleteData(r.Context(), userID, dataID)
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
