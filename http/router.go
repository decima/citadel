package http

import (
	"citadel-api/data/storage"
	"citadel-api/http/controller/access"
	"citadel-api/http/controller/blocks"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	blocksController := blocks.NewBlocksController(storage.NewBlockRepository())
	r.GET("/blocks", blocksController.GetTopLevel)
	r.GET("/blocks/:uuid", blocksController.Get)
	r.POST("/blocks", blocksController.Create)
	r.POST("/blocks/:parent", blocksController.Create)

	r.GET("/blocks/type/:type", blocksController.GetByType)

	accessController := access.NewAccessController(storage.NewBlockRepository())
	r.GET("/me", accessController.Me)
	// r.PUT("/blocks/:uuid", blocksController.Update)
	// r.DELETE("/blocks/:uuid", blocksController.Delete)

}
