package main

import (
	"log"

	s "github.com/fortify-presales/FortifyGoDemo/internal/server"
	h "github.com/fortify-presales/FortifyGoDemo/internal/echo"
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
	router := echo.New()

	router.GET("/ping/:cmd", func(c echo.Context) error {
		logger.Debugf("Received request: %s", c.Request().URL.Path)
		//
		// echo c.Param - Not yet supported by Fortify
		//
		host := c.Param("cmd")
		cmd := exec.Command("ping", "-c", "4", host)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return c.HTML(http.StatusInternalServerError, fmt.Sprintf("<pre>Error: %s</pre>", err))
		}
		return c.HTML(http.StatusOK, fmt.Sprintf("<pre>%s</pre>", output))
	})

	return router
}*/
