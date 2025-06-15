package main

import (
	"backend/cmd/api/routes"
	"os"
)

func main() {
	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
