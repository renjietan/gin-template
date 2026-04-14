package core

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GinLoggerMiddleware Gin中间件：记录每个HTTP请求的日志
func GinLoggerMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()
		fields := logrus.Fields{
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"query":      c.Request.URL.RawQuery,
			"status":     status,
			"client_ip":  c.ClientIP(),
			"latency_ms": latency.Milliseconds(),
			"user_agent": c.Request.UserAgent(),
		}
		if len(c.Errors) > 0 {
			fields["errors"] = c.Errors.String()
		}
		log.WithFields(fields).Info("HTTP request")
	}
}
