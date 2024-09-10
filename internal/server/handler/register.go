package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
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
//	@Success	200		{object}	model.UserResponse	"Response"
//	@Failure	409		{object}	model.Error			"Email is already used"
//	@Failure	500		{object}	model.Error			"Error"
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
	newUser, status, err := s.UserService.RegisterUser(r.Context(), user)
	if err != nil {
		s.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	s.JSON(w, http.StatusOK, model.UserResponse{
		ID:        newUser.ID,
		Email:     newUser.Email,
		CreatedAt: utils.Datetime(newUser.CreatedAt),
	})
}

// Login user
// Login godoc
//
//	@Summary	Login user
//	@Schemes
//	@Tags		Log
//	@Accept		json
//	@Produce	json
//	@Param		request	body		model.RegisterUser	true	"Body"
//	@Success	200		{object}	model.TokenResponse	"Response"
//	@Failure	409		{object}	model.Error			"Email is already used"
//	@Failure	500		{object}	model.Error			"Error"
//	@Router		/api/user/register [post]
func (s *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user model.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		s.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	if err := s.Validator.Struct(user); err != nil {
		s.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	token, status, err := s.UserService.Login(r.Context(), user)
	if err != nil {
		s.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	w.Header().Set("Authorization", token)
	s.SetTokenCookie(w, token, s.Config.Security.TokenExpSec)
	s.JSON(w, http.StatusOK, model.TokenResponse{
		Token: token,
	})
}
