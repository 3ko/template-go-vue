package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func AdminStatsHandler(c *gin.Context) {
    claims, _ := c.Get("claims")

    c.JSON(http.StatusOK, gin.H{
        "message": "Zone admin",
        "claims":  claims,
    })
}
