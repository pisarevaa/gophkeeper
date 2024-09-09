package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

// Register user
// RegisterUser godoc
//
//	@Summary	Regiser user
//	@Schemes
//	@Tags		Log
//	@Accept		json
//	@Produce	json
//	@Param		request	body		model.RegisterUser	true	"Body"
//	@Success	200		{object}	model.UserResponse		"Response"
//	@Failure	409		{object}	utils.Error			"Email is already used"
//	@Failure	500		{object}	utils.Error			"Error"
//	@Router		/api/user/register [post]
func (s *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		s.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	if err := s.Validator.Struct(user); err != nil {
		s.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	newUser, err := s.UserService.RegisterUser(r.Context(), user)
	if err != nil {
		s.JSON(w, http.StatusConflict, model.Error{Error: err.Error()})
		return
	}
	s.JSON(w, http.StatusOK, newUser)
}
