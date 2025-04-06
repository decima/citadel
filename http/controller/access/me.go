package access

import (
	"citadel-api/http/middlewares"
	"github.com/gin-gonic/gin"
)

func (*AccessController) Me(c *gin.Context) {
	user := middlewares.GetCurrentUser(c)
	c.JSON(200, user)
}
