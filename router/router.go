package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/achieveit-backend/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	applyCORS(router)

	authRoutes(router)

	return router
}

func applyCORS(router *gin.Engine) {
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"}
	router.Use(cors.New(config))
}

func authRoutes(router *gin.Engine) {
	authGroup := router.Group("/v1/auth")
	{
		authGroup.POST("/login", handlers.Login)
		authGroup.POST("/register", handlers.RegisterUser)
	}
}
