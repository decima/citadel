package middlewares

import (
	"citadel-api/data/services"
	"citadel-api/utils/container"
	"github.com/gin-gonic/gin"
)

const currentUserField = "currentUser"

func AuthHandler(c *gin.Context) {
	// get JWT from header
	jwt := c.Request.Header.Get("Authorization")
	if jwt == "" {
		c.JSON(401, gin.H{"error": "Missing JWT"})
		c.Abort()
		return
	}
	jwt = jwt[7:]
	// get Block from JWT
	user, err := container.ShouldGet[services.AccessManagerInterface]().GetFromJWT(jwt)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid JWT"})
		GetLogger(c).Error().Msgf("Invalid JWT: %v", err)
		c.Abort()
		return
	}

	c.Set(currentUserField, user)
	c.Next()

}

func GetCurrentUser(c *gin.Context) *services.User {
	block, exists := c.Get(currentUserField)
	if !exists {
		return nil
	}
	return block.(*services.User)
}
