package exception

import (
	"log"
	"net/http"
)

func Record(err error, w http.ResponseWriter, statusCode int) {
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(statusCode)
	return
}
