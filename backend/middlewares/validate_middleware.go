package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"exchangeapp/backend/validation"
)

// ValidationMiddleware is a validation middleware
func ValidationMiddleware(validate *validation.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj interface{}
		if err := c.BindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := validate.Validate(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
