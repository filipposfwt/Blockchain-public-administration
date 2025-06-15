package main

import (
	"log"

	"backend/internal/models"
	"backend/internal/staterecords"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	r := gin.Default()

	// Use the most basic CORS middleware
	r.Use(cors.Default())

	// Initialize state records handler
	stateRecordsHandler := staterecords.NewStateRecordsManager()

	// Routes
	r.POST("/documents", func(c *gin.Context) {
		var requestBody struct {
			StateInfo  models.StateInfo `json:"stateInfo"`
			PrivateKey string           `json:"privateKey"`
		}

		if err := c.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request format"})
			return
		}

		result, err := stateRecordsHandler.CreateStateRecord(requestBody.StateInfo, requestBody.PrivateKey)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, result)
	})

	r.GET("/documents/:recordId", func(c *gin.Context) {
		recordId := c.Param("recordId")
		hash := c.Query("hash")

		result, err := stateRecordsHandler.VerifyStateRecord(recordId, hash)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, result)
	})

	// Start server
	log.Println("Server starting on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
