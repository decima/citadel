package middlewares

import (
	"citadel-api/data/model"
	"citadel-api/data/services"
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
	block, err := services.NewAccessManager().GetFromJWT(jwt)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid JWT"})
		GetLogger(c).Error().Msgf("Invalid JWT: %v", err)
		c.Abort()
		return
	}

	c.Set(currentUserField, block)

}

func GetCurrentUser(c *gin.Context) *model.Block {
	block, exists := c.Get(currentUserField)
	if !exists {
		return nil
	}
	return block.(*model.Block)
}
