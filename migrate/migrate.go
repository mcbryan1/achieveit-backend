package main

import (
	"github.com/mcbryan1/achieveit-backend/initializers"
	"github.com/mcbryan1/achieveit-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Milestone{})
}
