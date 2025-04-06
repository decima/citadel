package blocks

import (
	"citadel-api/data/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (blocks *BlocksController) Create(c *gin.Context) {
	// Create a new block in the database
	block := &model.Block{}
	if err := c.BindJSON(block); err != nil {
		c.JSON(400, gin.H{"error": "Invalid block data"})
		return
	}
	if uuid := c.Param("parent"); uuid != "" {
		block.ParentID = &uuid
	}

	fmt.Println(block)

	// Validate the block data
	block2, err := blocks.repository.Create(block)
	fmt.Println(block2)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create block", "details": err.Error()})
		return
	}
	// Return the created block as JSON
	c.JSON(201, block)
}
