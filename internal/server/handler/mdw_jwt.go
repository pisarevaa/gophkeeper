package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

// Мидлвар по авторизации запросов по токену.
func (s *Handler) JWTAuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header["Authorization"]

		if len(authorization) != 1 {
			utils.JSON(w, http.StatusUnauthorized, model.Error{Error: "Authorization token is not set"})
			return
		}

		authHeader := authorization[0]
		if authHeader == "" {
			utils.JSON(w, http.StatusUnauthorized, model.Error{Error: "Authorization token is not set"})
			return
		}

		parts := strings.Split(authHeader, " ")
		var headersPartsCount = 2
		if len(parts) != headersPartsCount {
			utils.JSON(w, http.StatusUnauthorized, model.Error{Error: "Authorization token is not set"})
			return
		}

		token := parts[1]

		userID, err := utils.GetUserID(token, s.Config.Security.SecretKey)
		if err != nil {
			utils.JSON(w, http.StatusUnauthorized, model.Error{Error: "Authorization token is wrong"})
			return
		}
		ctx := context.WithValue(r.Context(), model.ContextUserID, userID)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
