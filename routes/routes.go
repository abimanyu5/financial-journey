package routes

import (
	"os"

	"financial-journey/controllers"
	"financial-journey/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up all the routes for the application
func SetupRouter(router *gin.Engine) {
	router.POST("/register", controllers.RegisterHandler)
	router.POST("/login", controllers.LoginHandler)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/user", controllers.GetUserHandler)
		protected.GET("/admin", controllers.AdminHandler)

		//transactions
		protected.POST("/transactions", controllers.InsertTransaction)
		protected.PUT("/transactions/:id", controllers.UpdateTransaction)
		protected.DELETE("/transactions/:id", controllers.DeleteTransaction)
		protected.GET("/transactions", controllers.GetAllTransactions)
		//master data
		protected.POST("/masters", controllers.InsertMaster)
		protected.PUT("/masters/:id", controllers.UpdateMaster)
		protected.DELETE("/masters/:id", controllers.DeleteMaster)
		protected.GET("/masters", controllers.GetAllMasters)
		protected.GET("/masters/:id/data", controllers.GetMasterById)

		//goals
		protected.POST("/goals", controllers.InsertGoals)
		protected.PUT("/goals/:id", controllers.UpdateGoals)
		protected.DELETE("/goals/:id", controllers.DeleteGoals)
		protected.GET("/goals", controllers.GetAllGoals)
		protected.GET("/goals/:id/data", controllers.GetGoalsById)
	}

	router.Run(":" + os.Getenv("PORT"))
}