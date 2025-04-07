package middlewares

import (
	"github.com/gin-gonic/gin"
)

func RoleHandler(expectedRole string) gin.HandlerFunc {

	return func(c *gin.Context) {
		user := GetCurrentUser(c)

		if user == nil {
			c.Abort()
			return
		}
		if user.Role != expectedRole {
			c.AbortWithStatusJSON(403, gin.H{
				"status": "Forbidden",
				"error":  "invalid role",
			})
			return
		}

	}

}
