package appkit

import (
	"errors"
	"fmt"
	"net/http"
)

// AppError – расширенная структура для кастомных ошибок.
// Помимо кода и сообщения, содержит "корневую" ошибку (Err) для детального логирования
// (в JSON не сериализуется).
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Error – реализует интерфейс error, чтобы AppError можно было использовать как обычную ошибку.
func (e *AppError) Error() string {
	if e.Err == nil {
		return e.Message
	}
	return fmt.Sprintf("%s | root error: %v", e.Message, e.Err)
}

// BadRequestError – 400 Bad Request
func BadRequestError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

// NotFoundError – 404 Not Found
func NotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// ValidationError – 422 Unprocessable Entity
func ValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

// InternalError – 500 Internal Server Error
func InternalError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func UnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func ForbiddenError(message string) *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

// WrapError – обертка над любой существующей error с дополнительным кодом и сообщением.
// Удобно, когда нужно сохранить контекст исходной ошибки.
func WrapError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// StatusCodeFromError – извлекает HTTP-код из error.
// Если это не AppError, возвращает 500 (Internal Server Error) по умолчанию.
func StatusCodeFromError(err error) int {
	if err == nil {
		return http.StatusOK
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code
	}

	return http.StatusInternalServerError
}
