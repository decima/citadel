package access

import "citadel-api/data/storage"

type AccessController struct {
	repository storage.BlockRepositoryInterface
}

func NewAccessController(repository storage.BlockRepositoryInterface) *AccessController {
	return &AccessController{
		repository: repository,
	}
}
