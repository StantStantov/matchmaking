package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			c.JSON(http.StatusInternalServerError, map[string]any{
				"Error":   err.Error(),
				"Success": false,
				"Value":   map[string]any{},
			})
		}
	}
}
