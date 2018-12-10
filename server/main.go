package server

import (
	"app.onca.api/server/config"
)

// Start ...
func Start() {
	config.Datastore()
	ConfigureRoutes()
}
