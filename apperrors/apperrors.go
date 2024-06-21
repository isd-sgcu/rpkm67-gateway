package apperrors

import "net/http"

type AppError struct {
	Id       string
	HttpCode int
}

func (e *AppError) Error() string {
	return e.Id
}

var (
	InternalServer     = &AppError{"Internal error", http.StatusInternalServerError}
	ServiceUnavailable = &AppError{"Internal error", http.StatusServiceUnavailable}
	Unauthorized       = &AppError{"Unauthorized", http.StatusUnauthorized}
	BadRequest         = &AppError{"Bad request", http.StatusBadRequest}
	InvalidToken       = &AppError{"Invalid token", http.StatusUnauthorized}
)

func BadRequestError(message string) *AppError {
	return &AppError{message, http.StatusBadRequest}
}

func NotFoundError(message string) *AppError {
	return &AppError{message, http.StatusNotFound}
}

func InternalServerError(message string) *AppError {
	return &AppError{message, http.StatusInternalServerError}
}
