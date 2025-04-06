package blocks

import "github.com/gin-gonic/gin"

func (blocks *BlocksController) Get(c *gin.Context) {
	uuid := c.Param("uuid")

	block, err := blocks.repository.Get(uuid)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get block"})
		return
	}

	// Return the block as JSON
	c.JSON(200, block)
}
