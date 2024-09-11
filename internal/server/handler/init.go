package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/pisarevaa/gophkeeper/internal/server/config"
	user "github.com/pisarevaa/gophkeeper/internal/server/service/auth"
)

type Handler struct {
	Config      config.Config
	Validator   *validator.Validate
	AuthService user.Service
}

type Option func(*Handler)

func WithConfig(config config.Config) Option {
	return func(s *Handler) {
		s.Config = config
	}
}

func WithValidator(validator *validator.Validate) Option {
	return func(s *Handler) {
		s.Validator = validator
	}
}

func WithAuthService(authService user.Service) Option {
	return func(s *Handler) {
		s.AuthService = authService
	}
}

// Создание хедлера роутера.
func NewHandler(options ...Option) *Handler {
	h := &Handler{}
	for _, option := range options {
		option(h)
	}
	return h
}

// Кодирование ответа в JSON.
func (h *Handler) JSON(w http.ResponseWriter, status int, model any) {
	bytes, err := json.Marshal(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status != http.StatusOK {
		slog.Error(string(bytes))
	}
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(bytes), status)
}

// Установка куки авторизации с токеном.
func (h *Handler) SetTokenCookie(w http.ResponseWriter, token string, tokenExpSec int64) {
	cookie := http.Cookie{}
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Duration(tokenExpSec))
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"
	http.SetCookie(w, &cookie)
}
