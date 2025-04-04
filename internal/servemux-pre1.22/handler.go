package handler

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler() http.Handler {
	router := http.NewServeMux()

	// GET request with data flow taint source in URL path
	router.HandleFunc("/api/v1/ping/{cmd}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling ping at %s\n", r.URL.Path)
		host := strings.TrimPrefix(r.URL.Path, "/api/v1/ping/")
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