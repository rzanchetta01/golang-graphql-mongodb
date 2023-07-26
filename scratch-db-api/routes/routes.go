package routes

import "scratch-db-api/controller"

func Startup() {
	controller.GetAllRecords()
}
