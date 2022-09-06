package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ministryofjustice/hosting-data-api/internal/tenants"
)

func main() {
	router := gin.Default()

	tenants.RegisterRoutes(router)

	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
