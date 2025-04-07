package access

import (
	"citadel-api/data/services"
	"citadel-api/utils/container"
	"github.com/gin-gonic/gin"
)

type AccessCreateRequest struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func (access *AccessController) Create(c *gin.Context) {
	accessRequest := AccessCreateRequest{}
	if err := c.ShouldBindJSON(&accessRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	token, err := container.ShouldGet[services.AccessManagerInterface]().Create(accessRequest.Name, accessRequest.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create access token"})
		return
	}

	c.JSON(200, gin.H{"token": token, "name": accessRequest.Name, "role": accessRequest.Role})
}
