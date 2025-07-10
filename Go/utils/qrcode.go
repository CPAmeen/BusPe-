package utils

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(upiID, busNumber string) error {
	url := fmt.Sprintf("https://buspe-payment-page.onrender.com/pay.html?upi=%s&bus=%s", upiID, busNumber)

	err := qrcode.WriteFile(url, qrcode.Medium, 256, fmt.Sprintf("qrcodes/%s.png", busNumber))
	if err != nil {
		return err
	}

	fmt.Println("✅ QR code generated for:", busNumber)
	return nil
}
