package Web

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type ServiceErrorDto struct {
	Message    string
	Err        error
	StatusCode int
}

func NewCustomServiceError(message string, err error, statusCode int) *ServiceErrorDto {
	return &ServiceErrorDto{Message: message, Err: err, StatusCode: statusCode}
}

func NewInternalServiceError(err error) *ServiceErrorDto {
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil
	}
	return NewCustomServiceError(http.StatusText(http.StatusInternalServerError), err, http.StatusInternalServerError)
}
