package http

import (
    "mon-projet/internal/config"
    "mon-projet/internal/http/handlers"
    "mon-projet/internal/http/middleware"
    "mon-projet/internal/repository"
    "mon-projet/internal/service"

    "github.com/gin-gonic/gin"
)

func NewRouter(cfg config.Config) *gin.Engine {
    r := gin.Default()

    db := repository.Connect()

    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    r.GET("/api/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    secure := r.Group("/api/secure")
    secure.Use(middleware.AuthZitadel())
    {
        secure.GET("/profile", handlers.ProfileHandler)
        secure.GET("/users", userHandler.GetAll)
        secure.POST("/users", userHandler.Create)
    }

    admin := r.Group("/api/secure/admin")
    admin.Use(middleware.AuthZitadel(), middleware.RequireRole("admin"))
    {
        admin.GET("/stats", handlers.AdminStatsHandler)
    }

    return r
}
