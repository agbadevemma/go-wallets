package middleware

import (
	"context"
	"uuid"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-wallets/internal/logger"
)

func CorrelationID() gin.HandlerFunc {
	return func(c *gin.Context) {
		corID := c.GetHeader("X-Correlation-ID")
		if corID == "" {
			corID = uuid.New().String()
		}
		ctx := context.WithValue(c.Request.Context(), logger.CorrelationIDKey, corID)
		c.Request = c.Request.WithContext(ctx)

		c.Header("x-Correlation-ID", corID)
		c.Next()
	}
}
