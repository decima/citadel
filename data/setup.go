package data

import (
	"citadel-api/data/services"
	"citadel-api/data/storage"
	"citadel-api/utils/logger"
)

func Setup() {
	repository := storage.NewBlockRepository()
	if counter, _ := repository.CountByType(services.UserBlock); counter > 0 {
		return
	}

	token, err := services.NewAccessManager().Create("admin", services.RoleAdmin)
	if err != nil {
		panic(err)
	}

	logger.Get().Warn().Msgf("Admin user created, access token: %s", token)

}
