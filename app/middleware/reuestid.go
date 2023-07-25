package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
)

const xRequestId = "X-Request-Id"

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, ues it if exists
		requestId := c.Request.Header.Get(xRequestId)

		// Create request id with UUID4
		if requestId == "" {
			requestId, _ = uuid.GenerateUUID()
		}

		// Expose it for use in the application
		c.Request = AddToContext(c, xRequestId, requestId)

		// Set X-Request-Id header
		c.Writer.Header().Set(xRequestId, requestId)
		c.Next()
	}
}
