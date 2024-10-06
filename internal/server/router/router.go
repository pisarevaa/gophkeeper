package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	r.Use(middleware.Logger)
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
	// Маршруты авторизации
	r.Post("/auth/register", handlers.RegisterUser)
	r.Post("/auth/login", handlers.Login)
	// Добавление маршрутов с авторизацией
	r.Mount("/api/data", AuthedRouter(handlers))
	return r
}

func AuthedRouter(handlers *handler.Handler) http.Handler {
	r := chi.NewRouter()
	r.Use(handlers.JWTAuthMiddleware)
	r.Get("/", handlers.GetData)
	r.Get("/{dataID}", handlers.GetDataByID)
	r.Post("/text", handlers.AddTextData)
	r.Post("/binary", handlers.AddBinaryData)
	r.Put("/text/{dataID}", handlers.UpdateTextData)
	r.Put("/binary/{dataID}", handlers.UpdateBinaryData)
	r.Delete("/{dataID}", handlers.DeleteData)
	return r
}
