package routes

import (
	"smart-command-center-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.Engine) {
	roleGroup := r.Group("/roles")
	{
		roleGroup.GET("/", controllers.GetAllRoles)
	}
}
