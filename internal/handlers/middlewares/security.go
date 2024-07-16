package middlewares

import (
	"net/http"
	"os"
	"strings"

	jwtUtils "github.com/Abdelrhmanfdl/user-service/internal/utils/jwt"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.Next()
			return
		}

		claims, err := jwtUtils.ValidateJWT(strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("userId", claims.Subject)
		ctx.Next()
	}
}

func AssertAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetString("userId")
		if userId == "" {
			ctx.Status(http.StatusUnauthorized)
		} else {
			ctx.Next()
		}
	}
}

func AssertUnauthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetString("userId")
		if userId == "" {
			ctx.Next()
		} else {
			ctx.Status(http.StatusForbidden)
		}
	}
}

func SecureGetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("X-API-Key") == os.Getenv("GET_USER_API_KEY") {
			ctx.Next()
		} else {
			ctx.Status(http.StatusForbidden)
		}
	}
}
