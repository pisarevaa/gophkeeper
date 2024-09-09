package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

type QueryMetrics struct {
	ID    string `json:"id"`
	MType string `json:"type"`
}

// Register user
// RegisterUser godoc
//
//	@Summary	Regiser user
//	@Schemes
//	@Tags		Log
//	@Accept		json
//	@Produce	json
//	@Param		request	body		model.RegisterUser	true	"Body"
//	@Success	200		{object}	model.Success		"Response"
//	@Failure	409		{object}	model.Error			"Email is already used"
//	@Failure	500		{object}	model.Error			"Error"
//	@Router		/api/user/register [post]
func (s *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	if err := s.Validator.Struct(user); err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.SuccessResponse(w)
}
