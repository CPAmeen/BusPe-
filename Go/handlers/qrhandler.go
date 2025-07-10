package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func GenerateQR(c *gin.Context) {
	upi := c.PostForm("upi")
	busNumber := c.PostForm("bus_number")

	if upi == "" || busNumber == "" {
		c.String(http.StatusBadRequest, "Missing UPI ID or Bus Number")
		return
	}

	qrURL := fmt.Sprintf("https://buspe-payment-page.onrender.com/pay.html?upi=%s&bus=%s", upi, busNumber)
	filename := fmt.Sprintf("%s_%s.png", upi, busNumber)
	filePath := filepath.Join("qrcodes", filename)

	err := qrcode.WriteFile(qrURL, qrcode.Medium, 256, filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate QR code")
		return
	}

	// Return HTML showing the QR code
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, fmt.Sprintf(`
		<html><body>
		<h3>QR Code for Bus Owner</h3>
		<img src="/static/%s" alt="QR Code"><br><br>
		<a href="/static/%s" download>Download QR</a>
		</body></html>`, filename, filename))
}
