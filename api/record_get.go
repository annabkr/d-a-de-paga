package api

import (
	"net/http"

	"github.com/annabkr/dia-de-paga/utils/errors"
)

func getRecord(w http.ResponseWriter, r *http.Request) error {
	return errors.NewForbiddenError("forbidden", nil, "123")
}
