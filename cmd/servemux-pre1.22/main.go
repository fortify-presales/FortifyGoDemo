package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/fortify-presales/FortifyGoDemo/internal/servemux-pre1.22/handler"
	
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      buildHandler(),
	}
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed\n")
	} else if err != nil {
		log.Printf("Error listening for server: %s\n", err)
	}
}
