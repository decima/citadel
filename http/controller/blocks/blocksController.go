package blocks

import (
	"citadel-api/data/storage"
)

type BlocksController struct {
	repository storage.BlockRepositoryInterface
}

func NewBlocksController(repository storage.BlockRepositoryInterface) *BlocksController {
	return &BlocksController{
		repository: repository,
	}
}
