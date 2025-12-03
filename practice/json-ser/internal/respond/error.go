package respond

import (
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Printf("Getting error: %v", err)
	}
	if code > 499 {
		log.Printf("Internal Server error, code: %v: %v\n", code, msg)
	}

	log.Printf("Getting error: %v\n", msg)
}
