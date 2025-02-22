package api

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Initialize Redis client
var rdb = redis.NewClient(&redis.Options{
    Addr: "localhost:6379", // Redis server address
    // Add other options like Password or DB if needed
})

// Logout handler
func (dbConfig *DBConfig) Logout(c *gin.Context) {
    // Extract refresh token from request (e.g., from Authorization header or cookie)
    refreshToken := c.GetHeader("Authorization")
    if refreshToken == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No refresh token provided"})
        return
    }

    // Invalidate the refresh token by removing it from Redis
    err := rdb.Del(ctx, refreshToken).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invalidate refresh token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
