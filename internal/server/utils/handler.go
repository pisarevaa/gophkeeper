package utils

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// Кодирование ответа в JSON.
func JSON(w http.ResponseWriter, status int, model any) {
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
func SetTokenCookie(w http.ResponseWriter, token string, tokenExpSec int64) {
	cookie := http.Cookie{}
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Duration(tokenExpSec))
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"
	http.SetCookie(w, &cookie)
}

// Получение ID данных из урра.
func GetDataID(r *http.Request) (int64, error) {
	dataIDString := chi.URLParam(r, "dataID")
	if dataIDString == "" {
		return 0, errors.New("path param dataID is not set")
	}
	dataID, err := strconv.ParseInt(dataIDString, 10, 64)
	if err != nil {
		return 0, errors.New("path param dataID is not integer")
	}
	return dataID, nil
}

// func GetDataID(w http.ResponseWriter, r *http.Request) {
// 	dataIDString := chi.URLParam(r, "dataID")
// 	if dataIDString == "" {
// 		JSON(w, http.StatusUnprocessableEntity, model.Error{Error: "Path param dataID is not set"})
// 		return
// 	}
// 	dataID, err := strconv.ParseInt(dataIDString, 10, 64)
//     if err != nil {
//         JSON(w, http.StatusUnprocessableEntity, model.Error{Error: "Path param dataID is not integer"})
// 		return
//     }
// }
