package middleware

import (
	// golang package
	"context"
	"order-service/infrastructure/log"
	"time"

	// external package
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// RequestLogger request logger.
//
// It returns gin.HandlerFunc when successful.
// Otherwise, empty gin.HandlerFunc will be returned.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		ctx := context.WithValue(timeoutCtx, "request_id", requestID)
		c.Request = c.Request.WithContext(ctx)

		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)

		requestLog := logrus.Fields{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(), // status code
			"latency":    latency,
		}

		if c.Writer.Status() == 200 || c.Writer.Status() == 201 {
			log.Logger.WithFields(requestLog).Info("Request Success")
		} else {
			log.Logger.WithFields(requestLog).Error("Request Error")
		}
	}
}
