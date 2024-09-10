package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	docs "github.com/pisarevaa/gophkeeper/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/pisarevaa/gophkeeper/internal/server/handler"
)

const MaxAge = 300

// Создание роутера.
func NewRouter(handlers *handler.Handler) chi.Router {
	// Инициализация роутера
	r := chi.NewRouter()
	// Логирование запросов
	r.Use(handlers.HTTPLoggingMiddleware)
	// Настройка корсов
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           MaxAge, // Maximum value not ignored by any of major browsers
	}))
	// Добавление сваггер документации
	docs.SwaggerInfo.BasePath = "/"
	r.Mount("/swagger", httpSwagger.WrapHandler)
	// if config.Key != "" {
	// 	r.Use(srv.HashCheckMiddleware)
	// }
	// Маршруты
	r.Post("/api/register", handlers.RegisterUser)
	r.Post("/api/login", handlers.Login)
	return r
}
