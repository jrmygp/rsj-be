package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"server/responses"
	userServices "server/services/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(userService userServices.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, responses.Response{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   "Token is Missing",
			})
			c.Abort()
			return

		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.Response{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   "Token is invalid",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.JSON(http.StatusUnauthorized, responses.Response{
					Code:   http.StatusUnauthorized,
					Status: "Unauthorized",
					Data:   "Token is expired",
				})
				c.Abort()
				return
			}

			// Fetch user data from the database using the repository
			userID := int(claims["user_id"].(float64)) // Convert user_id from float64
			user, err := userService.FindByID(userID)
			if err != nil {
				c.JSON(http.StatusUnauthorized, responses.Response{
					Code:   http.StatusUnauthorized,
					Status: "Unauthorized",
					Data:   "User with this token is not found!",
				})
				c.Abort()
				return
			}

			// Optionally set the user in the context
			c.Set("user", user)
			c.Next() // Call the next handler in the chain
		} else {
			c.JSON(http.StatusUnauthorized, responses.Response{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   "Token is invalid",
			})
			c.Abort()
			return
		}
	}
}
