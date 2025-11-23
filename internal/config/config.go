package config

import "os"

type Config struct {
	Port      string
	StaticDir string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "./client/dist"
	}

	return Config{
		Port:      port,
		StaticDir: staticDir,
	}
}
