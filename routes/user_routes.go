package routes

import (
	"smart-command-center-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.GET("/", controllers.GetUsers)
	users.GET("/:id", controllers.GetUser)
	users.POST("/", controllers.CreateUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
}
