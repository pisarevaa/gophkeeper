package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

// Register user
// RegisterUser godoc
//
//	@Summary	Regiser user
//	@Schemes
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		request	body		model.RegisterUser	true	"Body"
//	@Success	200		{object}	model.UserResponse	"Response"
//	@Failure	422		{object}	model.Error			"Unprocessable entity"
//	@Failure	409		{object}	model.Error			"Email is already used"
//	@Failure	500		{object}	model.Error			"Internal server error"
//	@Router		/auth/register [post]
func (s *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	if err := s.Validator.Struct(user); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	newUser, status, err := s.AuthService.RegisterUser(r.Context(), user)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	utils.JSON(w, http.StatusOK, model.UserResponse{
		ID:        newUser.ID,
		Email:     newUser.Email,
		CreatedAt: model.DateTime(newUser.CreatedAt),
	})
}

// Login user
// Login godoc
//
//	@Summary	Login user
//	@Schemes
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		request	body		model.RegisterUser	true	"Body"
//	@Success	200		{object}	model.TokenResponse	"Response"
//	@Failure	422		{object}	model.Error			"Unprocessable entity"
//	@Failure	404		{object}	model.Error			"Email is not found"
//	@Failure	401		{object}	model.Error			"Incorrect password"
//	@Failure	500		{object}	model.Error			"Internal server error"
//	@Router		/auth/login [post]
func (s *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user model.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	if err := s.Validator.Struct(user); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, model.Error{Error: err.Error()})
		return
	}
	token, status, err := s.AuthService.Login(r.Context(), user)
	if err != nil {
		utils.JSON(w, status, model.Error{Error: err.Error()})
		return
	}
	w.Header().Set("Authorization", token)
	utils.SetTokenCookie(w, token, s.Config.Security.TokenExpSec)
	utils.JSON(w, http.StatusOK, model.TokenResponse{
		Token: token,
		Email: user.Email,
	})
}
