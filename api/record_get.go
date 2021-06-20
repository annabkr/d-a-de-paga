package api

import (
	"net/http"

	"github.com/annabkr/paydayz/utils/errors"
)

func getRecord(w http.ResponseWriter, r *http.Request) error {
	return errors.NewForbiddenError("forbidden", nil, "123")
}
