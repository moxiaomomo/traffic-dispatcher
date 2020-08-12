package handler

import (
	"log"
	"net/http"
	"strings"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/admin") {
		log.Println(r.URL.Path)
	}
}
