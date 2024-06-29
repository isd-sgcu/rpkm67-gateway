package apperror

import "net/http"

type AppError struct {
	Id       string
	HttpCode int
}

func (e *AppError) Error() string {
	return e.Id
}

var (
	BadRequest         = &AppError{"Bad request", http.StatusBadRequest}
	NotFound           = &AppError{"Not found", http.StatusNotFound}
	InternalServer     = &AppError{"Internal error", http.StatusInternalServerError}
	Unauthorized       = &AppError{"Unauthorized", http.StatusUnauthorized}
	Forbidden          = &AppError{"Forbidden", http.StatusForbidden}
	ServiceUnavailable = &AppError{"Internal error", http.StatusServiceUnavailable}
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

func UnauthorizedError(message string) *AppError {
	return &AppError{message, http.StatusUnauthorized}
}

func ForbiddenError(message string) *AppError {
	return &AppError{message, http.StatusForbidden}
}

func ServiceUnavailableError(message string) *AppError {
	return &AppError{message, http.StatusServiceUnavailable}
}
