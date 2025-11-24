package api

import (
	"net/http"
	"path/filepath"
	"strings"

	"mon-projet/internal/api/handlers"
	"mon-projet/internal/api/middleware"
	"mon-projet/internal/config"
	"mon-projet/internal/db"
	"mon-projet/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg config.Config) *gin.Engine {
	r := gin.Default()

	database := db.Connect()

	userRepo := db.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	configService := service.NewConfigService(userService)
	configHandler := handlers.NewConfigHandler(configService)

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/api/config", configHandler.GetConfiguration)
	r.PUT("/api/config", configHandler.UpdateConfiguration)

	secure := r.Group("/api/secure")
	secure.Use(middleware.AuthZitadel())
	{
		secure.GET("/profile", handlers.ProfileHandler)
		secure.GET("/users", userHandler.GetAll)
		secure.POST("/users", userHandler.Create)
		secure.GET("/users/:id", userHandler.GetByID)
		secure.PUT("/users/:id", userHandler.Update)
		secure.DELETE("/users/:id", userHandler.Delete)
	}

	admin := r.Group("/api/secure/admin")
	admin.Use(middleware.AuthZitadel(), middleware.RequireRole("admin"))
	{
		admin.GET("/stats", handlers.AdminStatsHandler)
		admin.GET("/config", configHandler.GetConfiguration)
		admin.PUT("/config", configHandler.UpdateConfiguration)
	}

	if cfg.StaticDir != "" {
		staticIndex := filepath.Join(cfg.StaticDir, "index.html")

		r.Static("/assets", filepath.Join(cfg.StaticDir, "assets"))
		r.GET("/", func(c *gin.Context) {
			c.File(staticIndex)
		})

		r.NoRoute(func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
				return
			}

			c.File(staticIndex)
		})
	}

	return r
}
