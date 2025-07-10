package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckTokenMiddleware() gin.HandlerFunc {
	parser := jwt.NewParser()

	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, _, err := parser.ParseUnverified(tokenString, &userClaims{})
		if err != nil {
			c.Error(err)

			return
		}

		var userId string
		if claims, ok := token.Claims.(*userClaims); ok {
			userId = claims.UserId
		} else {
			c.Error(fmt.Errorf("CheckTokenMiddleware: Incorrect Token claims"))

			return
		}

		c.Set(tokenCtxKey, token)
		c.Set(userIdCtxKey, userId)

		c.Next()
	}
}

type userClaims struct {
	jwt.Claims
	UserId   string `json:"sub"`
	Username string `json:"username"`
	Role     string `json:"role"`
	IsActive string `json:"is_active"`
}

const userIdCtxKey = "UserId"

func GetUserId(c *gin.Context) string {
	id, ok := c.Get(userIdCtxKey)
	if !ok {
		return ""
	}

	return id.(string)
}

const tokenCtxKey = "Token"

func GetToken(c *gin.Context) *jwt.Token {
	token, ok := c.Get(tokenCtxKey)
	if !ok {
		return nil
	}

	return token.(*jwt.Token)
}
