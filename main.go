package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// func handleScrape(c echo.Context) error {

// 	return c.Attachment(f.name, f.name)
// }

// https://github.com/gin-gonic/gin#html-rendering
func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")
	// static file 두번째 파라미터에는 root path이기 때문에 ./이 필요하다.
	router.Static("/assets", "./assets")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run(":8080")
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
