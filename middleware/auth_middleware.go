package middleware

import (
	"manage_sales/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := c.GetHeader("Authorization")
        if tokenStr == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
            c.Abort()
            return
        }

        claims := &common.Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return common.JwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        if claims.ExpiresAt.Time.Before(time.Now()) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
            c.Abort()
            return
        }

        // Kiểm tra vai trò của người dùng
        roleIsValid := false
        for _, role := range requiredRoles {
            if claims.Role == role {
                roleIsValid = true
                break
            }
        }

        if !roleIsValid {
            c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
            c.Abort()
            return
        }

        c.Set("claims", claims)
        c.Next()
    }
}
