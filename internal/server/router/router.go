package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/pisarevaa/gophkeeper/internal/server/handler"
)

const MaxAge = 300

// Создание роутера.
func NewRouter(handlers *handler.Handler) chi.Router {
	// Инициализация роутера
	r := chi.NewRouter()
	// Логирование запросов
	r.Use(handlers.HTTPLoggingMiddleware)
	// Обработка запроса с сжатыми данными
	r.Use(handlers.GzipMiddleware)
	// Передача IP устройства
	r.Use(middleware.RealIP)
	// Настройка корсов
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           MaxAge, // Maximum value not ignored by any of major browsers
	}))

	// if config.Key != "" {
	// 	r.Use(srv.HashCheckMiddleware)
	// }
	// Маршруты
	r.Post("/api/logs/geo", handlers.SendGeoLog)
	return r
}
