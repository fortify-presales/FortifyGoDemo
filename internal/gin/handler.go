package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// BuildHandler sets up the HTTP routing and builds an HTTP handler.
func BuildHandler() http.Handler {
	router := gin.Default()

	// Create a sub router for the API Version
	v1 := router.Group("/api/v1")

	// GET request with data flow taint source in URL path
	v1.GET("/ping", func(c *gin.Context) {
		log.Printf("Handling GET at %s\n", c.Request.URL.Path)
		//
		// Get hostname from query parameter
		//
		// gin c.Query - Not yet supported by Fortify
		host := c.Query("hostname")
		if host == "" {
			c.String(http.StatusBadRequest, "Hostname not provided")
			return
		}
		//
		// Command Injection : dataflow
		//
		cmd := exec.Command("ping", "-c", "4", host)
		output, err := cmd.CombinedOutput()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("<pre>%s</pre>", output))
	})

	// POST request with data flow taint source in body
	v1.POST("/ping", func(c *gin.Context) {
		log.Printf("Handling POST at %s\n", c.Request.URL.Path)
		type JsonString struct {
			Hostname string `json:"hostname" binding:"required"`
		}
		var jsonDataToRead JsonString
		//
		//
		// gin c.ShouldBindJSON - Not yet supported by Fortify
		if err := c.ShouldBindJSON(&jsonDataToRead); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusOK, jsonDataToWrite)
	})

	// GET request with data flow taint source in URL path
	v1.GET("/download/:id", func(c *gin.Context) {
		log.Printf("Handling GET at %s\n", c.Request.URL.Path)
		//
		// Get id from URL path
		//
		// gin c.Param - Not yet supported by Fortify
		id := c.Param("id")
		if id == "" {
			c.String(http.StatusBadRequest, "Id not provided")
			return
		}
		//
		// Path Manipulation : dataflow
		//
		filename := fmt.Sprintf("%s%c%s%c%s", os.Getenv("PWD"), os.PathSeparator, "downloads", os.PathSeparator, id)
		log.Printf("Retrieving contents of file path: %s\n", filename)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			c.String(http.StatusNotFound, "File not found")
		}
		data, _ := ioutil.ReadFile(filename)
		c.JSON(http.StatusOK, string(data))
	})

	return router
}
