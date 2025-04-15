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