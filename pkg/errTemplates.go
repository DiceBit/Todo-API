package pkg

import (
	"context"
	"errors"
	"log"
	"net/http"
)

func ServerInternalError(err error, w http.ResponseWriter) bool {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return true
	}
	return false
}

func BadRequestError(err error, w http.ResponseWriter) bool {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return true
	}
	return false
}

func RequestTimeoutError(err error, w http.ResponseWriter) bool {
	if errors.Is(err, context.DeadlineExceeded) {
		http.Error(w, "Request timeout", http.StatusGatewayTimeout)
		return true
	}
	return false
}
