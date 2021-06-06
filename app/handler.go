package app

import (
	"fmt"
	"net/http"

	"github.com/annabkr/dia-de-paga/utils/errors"
	log "github.com/annabkr/dia-de-paga/utils/logger"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		switch e := err.(type) {
		case *errors.Error:
			log.Err(fmt.Sprintf("returning %d for request %s %s: %s", e.StatusCode(), r.Method, r.URL.String(), e.Error()))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(e.StatusCode())
			_, err = w.Write(e.Json())
			if err != nil {
				log.Warn(fmt.Sprintf("failed to write: %s", err.Error()))
			}
		default:
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}
	}
}
