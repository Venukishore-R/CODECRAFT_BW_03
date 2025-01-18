package routes

import (
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/auth"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, handlers *handlers.UserHandler, adminHandlers *handlers.AdminHandler) {
	api := r.Group("/api")
	{

		api.POST("/signup", handlers.SignUp)
		api.POST("/login", handlers.Login)
		{
			gUser := api.Group("/user")
			gUser.Use(auth.GeneralAuth())
			gUser.GET("/profile", handlers.Profile)
		}

		{
			gAdmin := api.Group("/admin")
			gAdmin.Use(auth.AdminAuth())
			gAdmin.GET("/users", adminHandlers.GetUsers)
			gAdmin.GET("/user", adminHandlers.GetUser)
			gAdmin.POST("/user", adminHandlers.Create)
			gAdmin.PUT("/user", adminHandlers.Update)
			gAdmin.DELETE("/user", adminHandlers.Delete)
		}
	}

}
