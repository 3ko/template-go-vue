package main

import (
	"log"
	apiserver "mon-projet/internal/api"
	"mon-projet/internal/config"
)

func main() {
	cfg := config.Load()

	r := apiserver.NewRouter(cfg)

	log.Printf("Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
