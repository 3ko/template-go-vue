package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileHandler(c *gin.Context) {
	claims, _ := c.Get("claims")
	roles, _ := c.Get("roles")

	c.JSON(http.StatusOK, gin.H{
		"claims": claims,
		"roles":  roles,
	})
}
