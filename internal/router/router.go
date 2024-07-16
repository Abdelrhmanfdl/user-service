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

	r.GET("/getUserData/:userId", middlewares.SecureGetUser(), routerHandler.HandleGetUserData)

	r.POST("/getUsersData/", middlewares.SecureGetUser(), routerHandler.HandleGetUsersData)

	r.Run(":" + os.Getenv("PORT"))
}
