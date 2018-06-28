package handler

import (
	"net/http"
)

func Headers(w http.ResponseWriter, contentType string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	switch contentType {
	case "text":
		w.Header().Set("Content-Type", "text/html")
	case "json":
		w.Header().Set("Content-Type", "application/json")
	}
}
