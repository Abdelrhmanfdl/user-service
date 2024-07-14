package router

import (
	"os"

	"github.com/Abdelrhmanfdl/user-service/internal/handlers"
	"github.com/Abdelrhmanfdl/user-service/internal/handlers/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(routerHandler *handlers.RouterHandler) {
	r := gin.Default()
	r.Use(middlewares.Authenticate())

	r.POST("/login", middlewares.AssertUnauthenticated(), routerHandler.HandleLogin)

	r.POST("/signup", middlewares.AssertUnauthenticated(), routerHandler.HandleSignup)

	// r.Run(":" + os.Getenv("PORT"))
	r.Run(":" + os.Args[1])
}
