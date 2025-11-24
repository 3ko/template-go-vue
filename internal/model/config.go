package model

// DatabaseConfig encapsulates connection details for a data source.
type DatabaseConfig struct {
	Host string `json:"host"`
	Name string `json:"name"`
	User string `json:"user"`
	Port string `json:"port"`
}

// AuthProviderConfig exposes identity provider settings such as Zitadel.
type AuthProviderConfig struct {
	Issuer      string `json:"issuer"`
	ClientID    string `json:"clientId"`
	RedirectURL string `json:"redirectUrl"`
}

// ApplicationConfig centralizes operational settings for the platform.
type ApplicationConfig struct {
	Database    DatabaseConfig     `json:"database"`
	ActiveUsers []User             `json:"activeUsers"`
	Auth        AuthProviderConfig `json:"auth"`
	Metadata    map[string]string  `json:"metadata,omitempty"`
	Configured  bool               `json:"configured"`
}
