package main

import (
	"github.com/CPAmeen/buspe/config"
	"github.com/CPAmeen/buspe/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnecttoDb()
	config.Sync()
}

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8300")
}
