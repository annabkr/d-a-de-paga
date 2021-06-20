package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/annabkr/paydayz/utils/logger"
)

type Error struct {
	Code       int    `json:"-"`
	Message    string `json:"message"`
	Underlying error  `json:"-"`
}

func (e Error) Error() string {
	return string(e.Json())
}

func (e Error) Json() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		log.Warn(fmt.Sprintf("unable to marshal json: %s", err.Error()))
	}
	return data
}

func (e Error) StatusCode() int {
	return e.Code
}

func NewError(message string, statusCode int, underlying error) error {
	return &Error{
		Message:    message,
		Code:       statusCode,
		Underlying: underlying,
	}
}

func NewBadRequestError(message string, underlying error) error {
	return NewError(message, http.StatusBadRequest, underlying)
}

func NewNotFoundError(message string, underlying error, requestId string) error {
	return NewError(message, http.StatusNotFound, underlying)
}

func NewForbiddenError(message string, underlying error, requestId string) error {
	return NewError(message, http.StatusForbidden, underlying)
}
