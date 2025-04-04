package main

import (
	"log"

	s "github.com/fortify-presales/FortifyGoDemo/internal/server"
	h "github.com/fortify-presales/FortifyGoDemo/internal/servemux-pre1.22"
	
)

func main() {
	srv := s.RunServer(8080, h.BuildHandler())
	if srv == nil {	
		log.Println("Server failed to start")
	}
}

/*
// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(logger log.Logger, cfg *config.Config) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/ping/{cmd}", func(w http.ResponseWriter, r *http.Request) {
		logger.Debugf("Received request: %s", r.URL.Path)
		//
		// gorilla mux.Vars - Not yet supported by Fortify
		//
		host := mux.Vars(r)["cmd"]
		cmd := exec.Command("ping", "-c", "4", host)
		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(output)
	})

	return router
}*/