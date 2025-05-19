package routes

import (
	"smart-command-center-backend/controllers"
	"smart-command-center-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	r.POST("/login", controllers.Login)

	authorized := r.Group("/api/v1")
	authorized.Use(middlewares.JWTAuthMiddleware())
	{
		// Role routes
		authorized.GET("/roles", controllers.GetAllRoles)

		// User routes
		authorized.GET("/users", controllers.GetUsers)
		authorized.GET("/users/:id", controllers.GetUser)
		authorized.POST("/users", controllers.CreateUser)
		authorized.PUT("/users/:id", controllers.UpdateUser)
		authorized.DELETE("/users/:id", controllers.DeleteUser)

	}

	return r
}
