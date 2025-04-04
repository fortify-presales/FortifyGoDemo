package handler

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// BuildHandler sets up the HTTP routing and builds an HTTP handler.
func BuildHandler() http.Handler {
	router := http.NewServeMux()

	// GET request with data flow taint source in URL path
	router.HandleFunc("GET /api/v1/ping/{hostname}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling GET at %s\n", r.URL.Path)
		// PathValue is new in Go 1.22
		host := r.PathValue("hostname")
		// Directly using user input in a shell command
		cmd := exec.Command("ping", "-c", "4", host)
		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
			return
		}
		// Return the command output to the user
		w.Header().Set("Content-Type", "text/plain")
		w.Write(output)
	})

	// GET request with data flow taint source in URL path
	router.HandleFunc("GET /api/v1/ping?hostname={hostname}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling GET at %s\n", r.URL.Path)
		// PathValue is new in Go 1.22
		host := r.PathValue("hostname")
		// Directly using user input in a shell command
		cmd := exec.Command("ping", "-c", "4", host)
		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
			return
		}
		// Return the command output to the user
		w.Header().Set("Content-Type", "text/plain")
		w.Write(output)
	})

	return router
}