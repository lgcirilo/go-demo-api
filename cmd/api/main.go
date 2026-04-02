package main

import (
	"github.com/gin-gonic/gin"

	"github.com/lgcirilo/go-demo-api/internal/config"
	"github.com/lgcirilo/go-demo-api/internal/db"
	"github.com/lgcirilo/go-demo-api/internal/handler"
	"github.com/lgcirilo/go-demo-api/internal/repository"
	"github.com/lgcirilo/go-demo-api/internal/service"
)

func main() {
	cfg := config.Load()

	dbPool := db.NewDBPool(cfg.DBUrl)

	userRepo := repository.NewUserRepository(dbPool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)

	r.Run(":" + cfg.Port)
}
