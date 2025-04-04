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
	router := gin.Default()

	router.GET("/ping/:cmd", func(c *gin.Context) {
		logger.Debugf("Received request: %s", c.Request.URL.Path)
		//
		// gin c.Param - Not yet supported by Fortify
		//
		host := c.Param("cmd")
		cmd := exec.Command("ping", "-c", "4", host)
		output, err := cmd.CombinedOutput()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err))
		}
		c.String(http.StatusOK, fmt.Sprintf("<pre>%s</pre>", output))
	})

	return router
}*/
