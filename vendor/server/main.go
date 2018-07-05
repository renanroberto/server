package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"server/router"
)

const PORT = "5000"

func main() {
	router := router.NewRouter()

	server := &http.Server{
		Addr:         getPORT(),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Server is running on http://localhost%s\n\n", getPORT())
	log.Fatal(server.ListenAndServe())
}

func getPORT() string {
	if port := os.Getenv("PORT"); len(port) > 0 {
		return ":" + port
	} else {
		return ":" + PORT
	}
}
