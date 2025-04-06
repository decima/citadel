package middlewares

import (
	"citadel-api/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"runtime/debug"
)

func LogHandler(c *gin.Context) {

	// Log the request
	l := logger.Get().With().
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Str("query", c.Request.URL.RawQuery).
		Str("client_ip", c.ClientIP()).
		Str("user_agent", c.Request.UserAgent()).
		Str("request_id", c.Request.Header.Get("X-Request-ID")).
		Logger()

	defer func() {

		if err := recover(); err != nil {
			l.Error().Msgf("Panic: %v", err)
			l.Debug().Msg(string(debug.Stack()))
			c.JSON(500, gin.H{"error": "Internal Server Error"})
		}
	}()

	c.Set("logger", l)
	c.Next()
	l.Info().Msgf("%v", c.Writer.Status())
}

func GetLogger(c *gin.Context) *zerolog.Logger {
	logTool, exists := c.Get("logger")

	if !exists {
		logTool = logger.Get()
		c.Set("logger", logTool)
	}

	l := logTool.(zerolog.Logger)

	return &l
}
