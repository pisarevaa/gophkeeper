package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pisarevaa/gophkeeper/internal/server/config"
	"github.com/pisarevaa/gophkeeper/internal/server/service/user"
)

type Handler struct {
	Config      config.Config
	Validator   *validator.Validate
	UserService user.Service
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

func WithUserService(userService user.Service) Option {
	return func(s *Handler) {
		s.UserService = userService
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

// Кодирование ответа в JSON
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
