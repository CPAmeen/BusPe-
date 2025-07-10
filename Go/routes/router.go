package routes

import (
	"github.com/CPAmeen/buspe/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all the application routes
func RegisterRoutes(router *gin.Engine) {
	// Static QR Code images
	router.Static("/static", "./qrcodes")

	// Load HTML templates
	router.LoadHTMLFiles("public/generate.html")

	// Show form
	router.GET("/generate", func(c *gin.Context) {
		c.HTML(200, "generate.html", nil)
	})
	//shows the success page
	router.LoadHTMLFiles("public/generate.html", "public/success.html")

	// Generate QR
	router.POST("/generate-qr", handlers.GenerateQR)

	// Transaction logging (if already created)
	router.POST("/log-transaction", handlers.LogTransaction)
}
