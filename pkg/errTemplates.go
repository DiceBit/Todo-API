package pkg

import (
	"log"
	"net/http"
)

func ServerInternalError(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return true
	}
	return false
}

func BadRequestError(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return true
	}
	return false
}
