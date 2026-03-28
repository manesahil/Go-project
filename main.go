package main

import (
	"Project/Routes"
	"Project/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB() // initialize DB

	router := gin.Default()

	Routes.UserRoutes(router)

	router.Run(":8080")
}
