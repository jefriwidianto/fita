package main

import (
	"fita/Config"
	"fita/Routes"
)

func main() {
	AppInitialization()
}

func AppInitialization() {
	//config DB SQL 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
	var setDB Config.ConfigSettingSql
	setDB.InitDB()

	//collect routes
	var routes Routes.Routes
	routes.CollectRoutes()
}
