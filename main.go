package main

import (
	"fita/Config"
	"fita/Routes"
)

func main() {
	AppInitialization()
}

func AppInitialization() {
	//config DB SQL test
	var setDB Config.ConfigSettingSql
	setDB.InitDB()

	//collect routes
	var routes Routes.Routes
	routes.CollectRoutes()
}
