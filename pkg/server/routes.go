package server

import "entry-notificator/internal/app/controllers"

// setRoutes sets routing
func (srv *server) setRoutes() {
	engine := srv.engine

	engine.GET("/ping", controllers.HealthCheckController)
}
