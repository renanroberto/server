package handler

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	Headers(w, "text")

	fmt.Fprintf(w, "Nothing here, but you can try /login and /auth")
}
