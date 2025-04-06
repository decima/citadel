package http

import (
	. "citadel-api/http/middlewares"
	"citadel-api/utils/build"
	"github.com/gin-gonic/gin"
)

func Start() error {
	// start a gin http on port 9009
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(LogHandler)
	r.Use(AuthHandler)
	r.GET("/", func(c *gin.Context) {
		c.String(200, `----------------
CITADEL API
----------------
Version: %s
Go Version: %s
`, build.Version(), build.GoVersion())
	})
	Route(r)

	return r.Run(":9009")
}
