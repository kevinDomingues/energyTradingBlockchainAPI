package middlewares

import (
	"energyTradingBlockchainAPI/pkg/services"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, BearerSchema) {
			c.AbortWithStatus(401)
		}

		tokenString := header[len(BearerSchema):]

		parsedToken, err := services.NewJWTService().ValidateToken(tokenString)
		if err != nil || !parsedToken.Valid {
			c.AbortWithStatus(401)
			return
		}

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			c.Set("claims", claims)
		} else {
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
