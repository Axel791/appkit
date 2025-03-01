package apikit

import (
	"encoding/json"
	"errors"
	"net/http"
)

// WriteJSON – универсальная функция для записи структуры в JSON с указанным статус-кодом.
func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

// WriteErrorJSON – удобная функция для вывода AppError (или обычной ошибки) в формате JSON.
func WriteErrorJSON(w http.ResponseWriter, err error) {
	// Извлекаем статус-код.
	code := StatusCodeFromError(err)

	var appErr *AppError
	if errors.As(err, &appErr) {
		WriteJSON(w, code, map[string]interface{}{
			"code":    code,
			"message": appErr.Message,
		})
		return
	}

	WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Internal Server Error",
	})
}
