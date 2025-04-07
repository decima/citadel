package data

import (
	"citadel-api/data/services"
	"citadel-api/data/storage"
	"citadel-api/utils/container"
	"citadel-api/utils/logger"
)

func Setup() {
	repository := container.ShouldGet[storage.BlockRepositoryInterface]()
	if counter, _ := repository.CountByType(services.UserBlock); counter > 0 {
		return
	}

	token, err := container.ShouldGet[services.AccessManagerInterface]().Create("admin", services.RoleAdmin)
	if err != nil {
		panic(err)
	}

	logger.Get().Warn().Msgf("Admin user created, access token: %s", token)

}
