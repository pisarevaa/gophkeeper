package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/utils"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

const maxUploadSize = 100 << 10

// Get all data
// GetData godoc
//
//	@Summary	Get all data
//	@Schemes
//	@Tags		Data
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]model.DataResponse	"Response"
//	@Failure	401	{object}	model.Error				"Unauthorized request"
//	@Failure	500	{object}	model.Error				"Internal server error"
//	@Router		/api/data [get]
func (s *Handler) GetData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(model.ContextUserID).(int64)
	if !ok {
		utils.JSON(w, http.StatusInternalServerError, model.Error{Error: "Error to cast userID into int64"})
		return
	}
	data, status, err := s.KeeperService.GetData(r.Context(), userID)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	var ordersResponse []model.DataResponseShort
	for _, d := range data {
		ordersResponse = append(
			ordersResponse,
			model.DataResponseShort{
				ID:   d.ID,
				Name: d.Name,
				Type: d.Type,
			},
		)
	}
	utils.JSON(w, http.StatusOK, ordersResponse)
}

// Get data by ID
// GetDataByID godoc
//
//	@Summary	Get data by ID
//	@Schemes
//	@Tags		Data
//	@Accept		json
//	@Produce	json
//	@Param		dataId			path		int					true	"Data ID"
//	@Param		Authorization	header		string				true	"Bearer"
//	@Success	200				{object}	model.DataResponse	"Response"
//	@Failure	422				{object}	model.Error			"Unprocessable entity (query)"
//	@Failure	404				{object}	model.Error			"Data is not found"
//	@Failure	401				{object}	model.Error			"Unauthorized request"
//	@Failure	500				{object}	model.Error			"Internal server error"
//	@Router		/api/data/{dataID} [get]
func (s *Handler) GetDataByID(w http.ResponseWriter, r *http.Request) { //nolint:dupl// it's okey
	userID, ok := r.Context().Value(model.ContextUserID).(int64)
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
		Name:      data.Name,
		Data:      data.Data,
		Type:      data.Type,
		CreatedAt: model.DateTime(data.CreatedAt),
		UpdatedAt: model.DateTime(data.UpdatedAt),
	})
}

// Add text data
// AddTextData godoc
//
//	@Summary	Add text data
//	@Schemes
//	@Tags		Data
//	@Accept		json
//	@Produce	json
//	@Param		request			body		model.AddTextData	true	"Body"
//	@Param		Authorization	header		string				true	"Bearer"
//	@Success	200				{object}	model.DataResponse	"Response"
//	@Failure	422				{object}	model.Error			"Unprocessable entity (body)"
//	@Failure	401				{object}	model.Error			"Unauthorized request"
//	@Failure	500				{object}	model.Error			"Internal server error"
//	@Router		/api/data/text [post]
func (s *Handler) AddTextData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(model.ContextUserID).(int64)
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
	data, status, err := s.KeeperService.AddTextData(r.Context(), textData.Name, textData.Data, userID)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	utils.JSON(w, http.StatusOK, model.DataResponse{
		ID:        data.ID,
		Name:      data.Name,
		Data:      data.Data,
		Type:      data.Type,
		CreatedAt: model.DateTime(data.CreatedAt),
		UpdatedAt: model.DateTime(data.UpdatedAt),
	})
}

// Add binary data
// AddTextData godoc
//
//	@Summary	Add text data
//	@Schemes
//	@Tags		Data
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		file			formData	file				true	"File"
//	@Param		name			formData	string				true	"Name"
//	@Param		Authorization	header		string				true	"Bearer"
//	@Success	200				{object}	model.DataResponse	"Response"
//	@Failure	422				{object}	model.Error			"Unprocessable entity (body)"
//	@Failure	401				{object}	model.Error			"Unauthorized request"
//	@Failure	500				{object}	model.Error			"Internal server error"
//	@Router		/api/data/binary [post]
func (s *Handler) AddBinaryData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(model.ContextUserID).(int64)
	if !ok {
		utils.JSON(w, http.StatusInternalServerError, model.Error{Error: "Error to cast userID into int64"})
		return
	}
	// Читаем данные не больше 10 Мб
	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	// Получение имени
	name := r.FormValue("name")
	if name == "" {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: "Name is empty"})
		return
	}
	file, err := utils.ReadBinaryData(r)
	if err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	data, status, err := s.KeeperService.AddBinaryData(r.Context(), file, name, userID)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	utils.JSON(w, http.StatusOK, model.DataResponse{
		ID:        data.ID,
		Name:      data.Name,
		Data:      data.Data,
		Type:      data.Type,
		CreatedAt: model.DateTime(data.CreatedAt),
		UpdatedAt: model.DateTime(data.UpdatedAt),
	})
}

// Update text data
// UpdateTextData godoc
//
//	@Summary	Update text data
//	@Schemes
//	@Tags		Data
//	@Accept		json
//	@Produce	json
//	@Param		dataId			path		int					true	"Data ID"
//	@Param		request			body		model.AddTextData	true	"Body"
//	@Param		Authorization	header		string				true	"Bearer"
//	@Success	200				{object}	model.DataResponse	"Response"
//	@Failure	422				{object}	model.Error			"Unprocessable entity (query or body)"
//	@Failure	401				{object}	model.Error			"Unauthorized request"
//	@Failure	500				{object}	model.Error			"Internal server error"
//	@Router		/api/data/text/{dataID} [put]
func (s *Handler) UpdateTextData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(model.ContextUserID).(int64)
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
	if errDecode := json.NewDecoder(r.Body).Decode(&textData); errDecode != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: errDecode.Error()})
		return
	}
	if errValidate := s.Validator.Struct(textData); errValidate != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: errValidate.Error()})
		return
	}
	data, status, err := s.KeeperService.UpdateTextData(r.Context(), textData.Name, textData.Data, userID, dataID)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	utils.JSON(w, http.StatusOK, model.DataResponse{
		ID:        data.ID,
		Name:      data.Name,
		Data:      data.Data,
		Type:      data.Type,
		CreatedAt: model.DateTime(data.CreatedAt),
		UpdatedAt: model.DateTime(data.UpdatedAt),
	})
}

// Update binary data
// UpdateBinaryData godoc
//
//	@Summary	Update binary data
//	@Schemes
//	@Tags		Data
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		dataId			path		int					true	"Data ID"
//	@Param		file			formData	file				true	"File"
//	@Param		name			formData	string				true	"Name"
//	@Param		Authorization	header		string				true	"Bearer"
//	@Success	200				{object}	model.DataResponse	"Response"
//	@Failure	422				{object}	model.Error			"Unprocessable entity (query or body)"
//	@Failure	401				{object}	model.Error			"Unauthorized request"
//	@Failure	500				{object}	model.Error			"Internal server error"
//	@Router		/api/data/binary/{dataID} [put]
func (s *Handler) UpdateBinaryData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(model.ContextUserID).(int64)
	if !ok {
		utils.JSON(w, http.StatusInternalServerError, model.Error{Error: "Error to cast userID into int64"})
		return
	}
	dataID, err := utils.GetDataID(r)
	if err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	// Читаем данные не больше 10 Мб
	err = r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	// Получение имени
	name := r.FormValue("name")
	if name == "" {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: "Name is empty"})
		return
	}
	file, err := utils.ReadBinaryData(r)
	if err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	data, status, err := s.KeeperService.UpdateBinaryData(r.Context(), name, file, userID, dataID)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	utils.JSON(w, http.StatusOK, model.DataResponse{
		ID:        data.ID,
		Name:      data.Name,
		Data:      data.Data,
		Type:      data.Type,
		CreatedAt: model.DateTime(data.CreatedAt),
		UpdatedAt: model.DateTime(data.UpdatedAt),
	})
}

// Delete data
// DeleteData godoc
//
//	@Summary	Delete data
//	@Schemes
//	@Tags		Data
//	@Accept		json
//	@Produce	json
//	@Param		dataId			path		int					true	"Data ID"
//	@Param		Authorization	header		string				true	"Bearer"
//	@Success	200				{object}	model.DataResponse	"Response"
//	@Failure	422				{object}	model.Error			"Unprocessable entity (query)"
//	@Failure	401				{object}	model.Error			"Unauthorized request"
//	@Failure	500				{object}	model.Error			"Internal server error"
//	@Router		/api/data/{dataID} [delete]
func (s *Handler) DeleteData(w http.ResponseWriter, r *http.Request) { //nolint:dupl// it's okey
	userID, ok := r.Context().Value(model.ContextUserID).(int64)
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
		Name:      data.Name,
		Data:      data.Data,
		Type:      data.Type,
		CreatedAt: model.DateTime(data.CreatedAt),
		UpdatedAt: model.DateTime(data.UpdatedAt),
	})
}
