package server

import "entry-notificator/internal/app/controllers"

// setRoutes sets routing
func (srv *server) setRoutes() {
	engine := srv.engine

	engine.GET("/ping", controllers.HealthCheckController)

	engine.POST("/children", controllers.CreateChild)
	engine.POST("/cards", controllers.CreateCard)
	engine.POST("/places/:place_id/attendances", controllers.CreateAttendance)
	engine.GET("/attendances", controllers.IndexAttendances)
	engine.POST("/supporters", controllers.CreateSupporter)
}
