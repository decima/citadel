package blocks

import "github.com/gin-gonic/gin"

func (blocks *BlocksController) GetTopLevel(c *gin.Context) {
	// Get all blocks from the database
	blocksList, err := blocks.repository.GetChildren(nil)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get blocks"})
		return
	}

	// Return the blocks as JSON
	c.JSON(200, blocksList)
}
