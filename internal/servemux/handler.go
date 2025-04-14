package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// BuildHandler sets up the HTTP routing and builds an HTTP handler.
func BuildHandler() http.Handler {
	router := http.NewServeMux()

	// GET request with data flow taint source in URL path
	router.HandleFunc("GET /api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling GET at %s\n", r.URL.Path)
		//
		// Get hostname from query parameter
		//
		host := r.URL.Query().Get("hostname")
		if host == "" {
			http.Error(w, "Hostname not provided", http.StatusBadRequest)
			return
		}
		//
		// Command Injection : dataflow
		//
		cmd := exec.Command("ping", "-c", "4", host)
		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(output)
	})

	// POST request with data flow taint source in body
	router.HandleFunc("POST /api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling POST at %s\n", r.URL.Path)
		type JsonString struct {
			Hostname string `json:"hostname"`
		}
		var jsonDataToRead JsonString
		//err2 := json.Unmarshal(body, &jsonData)
		err := json.NewDecoder(r.Body).Decode(&jsonDataToRead)
		if err != nil {
			http.Error(w, "Hostname not provided", http.StatusBadRequest)
			return
		}
		//
		// JSON Injection : dataflow
		//
		jsonDataToWrite := map[string]string{
			"command":  "ping",
			"hostname": jsonDataToRead.Hostname,
			"output":   "", // Placeholder for actual output
		}
		log.Printf("Creating file 'command_log.json' with contents: %+v\n", jsonDataToWrite)
		file, _ := os.OpenFile("command_log.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer file.Close()
		jsonEncoder := json.NewEncoder(file)
		jsonEncoder.SetIndent("", "  ") // Optional: Pretty-print the JSON
		jsonEncoder.Encode(jsonDataToWrite)
		// TODO: actual ping logic can be added here and output placed in "jsonDataToWrite"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	// GET request with data flow taint source in URL path
	router.HandleFunc("GET /api/v1/download/{id}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling GET at %s\n", r.URL.Path)
		//
		// Get id from URL path
		//
		// PathValue is new in Go 1.22 - Not yet supported by Fortify
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Id not provided", http.StatusBadRequest)
			return
		}
		//
		// Path Manipulation : dataflow
		//
		filename := fmt.Sprintf("%s%c%s%c%s", os.Getenv("PWD"), os.PathSeparator, "downloads", os.PathSeparator, id)
		log.Printf("Retrieving contents of file path: %s\n", filename)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		data, _ := ioutil.ReadFile(filename)
		w.Header().Set("Content-Type", "text/plain")
		w.Write(data)
	})

	return router
}
