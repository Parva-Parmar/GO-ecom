// package middleware

// import(
// 	token "github.com/Parva-Parmar/GO-ecom/tokens"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// func Authentication() gin.HandlerFunc{
// 	return func(c *gin.Context){
// 		ClientToken := c.Request.Header.Get("token")
// 		if ClientToken == ""{
// 			c.JSON(http.StatusInternalServerError,gin.H{"error":"No authorization header provided"})
// 			c.Abort()
// 			return
// 		}
// 		claims,err := token.ValidateToken(ClientToken)
// 		if err != ""{
// 			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("email",claims.Email)
// 		c.Set("uid",claims.Uid)
// 		c.Next()
// 	}
// }

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/Parva-Parmar/GO-ecom/tokens" // adjust the import path as necessary
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the Authorization header
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}

		// Split the header to check for the Bearer scheme
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// Extract the token part
		clientToken := parts[1]
		claims, err := token.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		// Set claims in the context for further use in handlers
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
