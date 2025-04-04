package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// BuildHandler sets up the HTTP routing and builds an HTTP handler.
func BuildHandler() http.Handler {
	router := http.NewServeMux()

	// GET request with data flow taint source in URL query parameter
	// POST request with data flow taint source in body
	router.HandleFunc("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
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
		} else if r.Method == http.MethodPost {
			log.Printf("Handling POST at %s\n", r.URL.Path)
			defer r.Body.Close()
			body, err1 := ioutil.ReadAll(r.Body)
			if err1 != nil{
				http.Error(w, "No body provided", http.StatusBadRequest)
				return
			}
			type JsonString struct {
				Hostname string `json:"hostname"`
			}
			var jsonData JsonString
			err2 := json.Unmarshal(body, &jsonData)
			if err2 != nil {
				http.Error(w, "Hostname not provided", http.StatusBadRequest)
				return
			}
			//
			// JSON Injection : dataflow
			//
			jsonString := `{
				"command":"` + "ping" + `",
				"hostname":"` + jsonData.Hostname + `",
			}`
			file, _ := os.OpenFile("command_log.json", os.O_CREATE, os.ModePerm) 
			defer file.Close()  
			jsonEncoder := json.NewEncoder(file) 
			jsonEncoder.Encode(jsonString)
			// TODO: actual ping logic can be added here
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	// GET request with data flow taint source in URL path
	router.HandleFunc("/api/v1/download/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			log.Printf("Handling GET at %s\n", r.URL.Path)
			//
			// Get id from URL path
			//
			id := strings.TrimPrefix(r.URL.Path, "/api/v1/download/")
			if id == "" {	
				http.Error(w, "Id not provided", http.StatusBadRequest)
				return
			}
			//
			// Path Manipulation : dataflow
			//
			filename := fmt.Sprintf("%s%c%s%c%s", os.Getenv("PWD"), os.PathSeparator, "etc", os.PathSeparator, id)
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				http.Error(w, "File not found", http.StatusNotFound)
				return
			}
    		data, _ := ioutil.ReadFile(filename)
			w.Header().Set("Content-Type", "text/plain")
			w.Write(data)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	return router
}