package handlers

import (
	"net/http"
	"time"

	"github.com/CPAmeen/buspe/config"
	"github.com/CPAmeen/buspe/models"
	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	FareAmount float64   `json:"fare_amount"`
	BusNumber  string    `json:"bus_number"`
	UPI        string    `json:"upi_id"`
	PaidAt     time.Time `json:"paid_at"`
}

func LogTransaction(c *gin.Context) {
	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	platformFee := input.FareAmount * 0.10

	tx := models.Transaction{
		BusID:       input.BusNumber,
		FareAmount:  input.FareAmount,
		PlatformFee: platformFee,
		PaymentMode: "digital",
		PaidAt:      input.PaidAt,
	}

	if err := config.DB.Create(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "transaction_id": tx.ID})
}
