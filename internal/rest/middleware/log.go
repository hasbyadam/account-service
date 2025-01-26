package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func RequestLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		res := next(c)

		log.WithFields(log.Fields{
			"method":     c.Request().Method,
			"path":       c.Path(),
			"user_agent": c.Request().UserAgent(),
			"latency_ns": time.Since(start).Nanoseconds(),
		}).Info("incoming request")

		return res
	}
}
