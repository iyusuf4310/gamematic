package errs

import (
	"k/golang/gamematic/logger"
	"net/http"
)

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func NewNotFoundError(message string) *AppError {
	logger.Error(message)
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	logger.Error(message)
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	logger.Error(message)
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
