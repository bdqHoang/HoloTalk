package router

import (
	"api-gateway/handlers"

	"github.com/gin-gonic/gin"
)

func RouterConfig() *gin.Engine {
	r := gin.Default()
	
	authRouter := r.Group("/auth")
	{
		authRouter.POST("/register", handlers.RegisterHandler)
		authRouter.POST("/login", handlers.LoginHandler)
		authRouter.POST("/refresh", handlers.RefreshTokenHandler)
		//authRouter.POST("/logout", handlers.LogoutHandler)
	}
	
	return r

}