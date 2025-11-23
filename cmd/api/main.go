package main

import (
    "log"
    "mon-projet/internal/config"
    httpserver "mon-projet/internal/http"
)

func main() {
    cfg := config.Load()

    r := httpserver.NewRouter(cfg)

    log.Printf("Server running on port %s", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatal(err)
    }
}
