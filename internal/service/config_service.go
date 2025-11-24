package service

import (
	"os"
	"sync"

	"mon-projet/internal/model"
)

// ConfigService centralizes runtime configuration and aggregates operational insights.
type ConfigService struct {
	mu          sync.RWMutex
	config      model.ApplicationConfig
	userService *UserService
}

// NewConfigService builds a configuration service with sane defaults hydrated from the environment.
func NewConfigService(userService *UserService) *ConfigService {
	cfg := model.ApplicationConfig{
		Database: model.DatabaseConfig{
			Host: envOrDefault("DB_HOST", "localhost"),
			Name: envOrDefault("DB_NAME", "app"),
			User: envOrDefault("DB_USER", "admin"),
			Port: envOrDefault("DB_PORT", "5432"),
		},
		Auth: model.AuthProviderConfig{
			Issuer:      envOrDefault("ZITADEL_ISSUER", "https://example.zitadel.cloud"),
			ClientID:    envOrDefault("ZITADEL_CLIENT_ID", "client-id"),
			RedirectURL: envOrDefault("ZITADEL_REDIRECT", "http://localhost:5173/callback"),
		},
		Metadata:   map[string]string{},
		Configured: false,
	}

	return &ConfigService{
		config:      cfg,
		userService: userService,
	}
}

// GetConfiguration returns the current configuration along with dynamic details such as active users.
func (s *ConfigService) GetConfiguration() (model.ApplicationConfig, error) {
	s.mu.RLock()
	base := s.config
	s.mu.RUnlock()

	if s.userService == nil {
		return base, nil
	}

	users, err := s.userService.GetAll()
	if err != nil {
		return model.ApplicationConfig{}, err
	}

	base.ActiveUsers = users
	return base, nil
}

// UpdateConfiguration applies incoming configuration updates, preserving unspecified fields.
func (s *ConfigService) UpdateConfiguration(update model.ApplicationConfig) (model.ApplicationConfig, error) {
	s.mu.Lock()
	if update.Database.Host != "" {
		s.config.Database.Host = update.Database.Host
	}
	if update.Database.Name != "" {
		s.config.Database.Name = update.Database.Name
	}
	if update.Database.User != "" {
		s.config.Database.User = update.Database.User
	}
	if update.Database.Port != "" {
		s.config.Database.Port = update.Database.Port
	}
	if update.Metadata != nil {
		s.config.Metadata = update.Metadata
	}
	if update.Auth.Issuer != "" {
		s.config.Auth.Issuer = update.Auth.Issuer
	}
	if update.Auth.ClientID != "" {
		s.config.Auth.ClientID = update.Auth.ClientID
	}
	if update.Auth.RedirectURL != "" {
		s.config.Auth.RedirectURL = update.Auth.RedirectURL
	}
	s.config.Configured = true
	s.mu.Unlock()

	return s.GetConfiguration()
}

func envOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
