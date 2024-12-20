package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/achieveit-backend/handlers"
	"github.com/mcbryan1/achieveit-backend/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	applyCORS(router)

	authRoutes(router)
	goalRoutes(router)
	milestoneRoutes(router)

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

func goalRoutes(router *gin.Engine) {
	goalGroup := router.Group("/v1/goals")
	goalGroup.Use(middlewares.AuthMiddleware())
	{
		goalGroup.POST("/create-goal", handlers.CreateGoal)
		goalGroup.GET("/fetch-goals", handlers.GetGoals)
		goalGroup.GET("/fetch-goal/:id", handlers.GetGoal)
		goalGroup.PUT("/update-goal/:id", handlers.UpdateGoal)
		goalGroup.DELETE("/delete-goal/:id", handlers.DeleteGoal)
	}

}

func milestoneRoutes(router *gin.Engine) {
	milestoneGroup := router.Group("/v1/milestones")
	milestoneGroup.Use(middlewares.AuthMiddleware())
	{
		milestoneGroup.POST("/create-milestone", handlers.CreateMilestone)
		milestoneGroup.GET("/fetch-milestones", handlers.GetMilestones)
		// milestoneGroup.GET("/fetch-milestone/:id", handlers.GetMilestone)
		// milestoneGroup.PUT("/update-milestone/:id", handlers.UpdateMilestone)
		// milestoneGroup.DELETE("/delete-milestone/:id", handlers.DeleteMilestone)
	}

}
