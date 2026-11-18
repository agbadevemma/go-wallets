package middleware

import (
	"github.com/gin-gonic/gin"
	customError "github.com/yourusername/go-wallets/internal/errors"
	"github.com/yourusername/go-wallets/internal/logger"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*customError.AppError); ok {
				logger.Warn(c.Request.Context(), "Client error occured",
					"code", appErr.Code,
					"message", appErr.Message,
					"status", appErr.StatusCode,
				)
				c.JSON(appErr.StatusCode, gin.H{
					"success": false,
					"error":   appErr,
				})
				return
			}

		}

	}
}
