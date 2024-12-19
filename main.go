package main

import (
	"fmt"
	"os"

	"github.com/mcbryan1/achieveit-backend/initializers"
	"github.com/mcbryan1/achieveit-backend/router"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := router.SetupRouter()

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("PORT environment variable not set")
	}

	router.Run(":" + port)
}
