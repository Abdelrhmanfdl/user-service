package main

import (
	"github.com/Abdelrhmanfdl/user-service/internal/handlers"
	"github.com/Abdelrhmanfdl/user-service/internal/router"
	"github.com/Abdelrhmanfdl/user-service/internal/service"
)

func main() {
	userService := service.NewUserService()
	handler := handlers.NewRouterHandler(userService)
	router.InitRouter(handler)

}
