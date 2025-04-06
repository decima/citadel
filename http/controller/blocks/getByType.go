package blocks

import "github.com/gin-gonic/gin"

func (blocks *BlocksController) GetByType(context *gin.Context) {
	typeName := context.Param("type")

	blocksList, err := blocks.repository.GetByType(typeName)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to get blocks"})
		return
	}

	// Return the blocks as JSON
	context.JSON(200, blocksList)

}
