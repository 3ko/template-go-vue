package middleware

import (
	"net/http"
	"os"
	"strings"

	"mon-projet/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type RoleSet map[string]bool

func AuthZitadel() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Bearer token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		issuer := os.Getenv("ZITADEL_ISSUER")
		audience := os.Getenv("ZITADEL_AUDIENCE")

		if issuer == "" || audience == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Zitadel env not configured"})
			c.Abort()
			return
		}

		jwksURL := issuer + "/oauth/v2/keys"
		jwks, err := utils.LoadJWKS(jwksURL)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed loading JWKS"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, jwks.Keyfunc)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			c.Abort()
			return
		}

		if iss, ok := claims["iss"].(string); !ok || iss != issuer {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid issuer"})
			c.Abort()
			return
		}

		if !containsAudience(claims["aud"], audience) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid audience"})
			c.Abort()
			return
		}

		roles := extractRolesFromClaims(claims)

		c.Set("claims", claims)
		c.Set("roles", roles)

		c.Next()
	}
}

func containsAudience(aud interface{}, target string) bool {
	switch v := aud.(type) {
	case string:
		return v == target
	case []interface{}:
		for _, a := range v {
			if s, ok := a.(string); ok && s == target {
				return true
			}
		}
	}
	return false
}

func extractRolesFromClaims(claims jwt.MapClaims) RoleSet {
	rs := RoleSet{}

	if v, ok := claims["urn:zitadel:iam:roles"]; ok {
		addRoles(rs, v)
	}
	if v, ok := claims["roles"]; ok {
		addRoles(rs, v)
	}

	return rs
}

func addRoles(rs RoleSet, raw interface{}) {
	switch v := raw.(type) {
	case []interface{}:
		for _, r := range v {
			if s, ok := r.(string); ok {
				rs[s] = true
			}
		}
	case map[string]interface{}:
		for k, val := range v {
			if b, ok := val.(bool); ok && b {
				rs[k] = true
			}
		}
	case string:
		rs[v] = true
	}
}
