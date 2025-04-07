package http

import (
	"citadel-api/data/services"
	"citadel-api/data/storage"
	"citadel-api/http/controller/access"
	"citadel-api/http/controller/blocks"
	"citadel-api/http/middlewares"
	"citadel-api/utils/container"
	"github.com/gin-gonic/gin"
)

func init() {
}

func Route(r *gin.Engine) {
	blockRepository := container.ShouldGet[storage.BlockRepositoryInterface]()

	blocksController := blocks.NewBlocksController(blockRepository)
	r.GET("/blocks", blocksController.GetTopLevel)
	r.GET("/blocks/:uuid", blocksController.Get)
	r.POST("/blocks", blocksController.Create)
	r.POST("/blocks/:parent", blocksController.Create)

	r.GET("/blocks/type/:type", blocksController.GetByType)

	accessController := access.NewAccessController(blockRepository)
	r.GET("/me", middlewares.RoleHandler(services.RoleAdmin), accessController.Me)
	// r.PUT("/blocks/:uuid", blocksController.Update)
	// r.DELETE("/blocks/:uuid", blocksController.Delete)

}
