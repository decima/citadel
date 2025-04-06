package main

import (
	"citadel-api/data"
	"citadel-api/http"
	"citadel-api/utils/build"
	"citadel-api/utils/logger"
)

var Version string = "dev"

func main() {
	build.Initialize(Version)
	data.Setup()
	webapp := http.Start()
	logger.Get().Warn().Err(webapp).Msg("Server stopped")

}
