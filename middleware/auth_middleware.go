
package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequiredAuth() gin.HandlerFunc {
	return func(context *gin.Context) {

		// 1. Mengambil token dari header
		authHeader := context.GetHeader("Authorization")
		parts := strings.Fields(authHeader)

		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Akses ditolak, token tidak ditemukan",
			})
			return
		}

		tokenString := parts[1]

		// 2. Mengambil JWT Secret
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "JWT_SECRET tidak ditemukan",
			})
			return
		}

		// 3. Memvalidasi token
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("signing method tidak valid")
				}

				return []byte(secret), nil
			},
		)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Token tidak valid",
				"detail": err.Error(),
			})
			return
		}

		if !token.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token tidak valid",
			})
			return
		}

		// ambil user id
		sub := fmt.Sprint(claims["sub"])
		userID, err := strconv.Atoi(sub)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "User ID tidak valid",
			})
			return
		}

		// 5. Menyimpan User ID ke context
		context.Set("userId", userID)

		context.Next()
	}
}
