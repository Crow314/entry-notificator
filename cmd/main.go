package main

import (
	"entry-notificator/pkg/config"
	"entry-notificator/pkg/database"
	"entry-notificator/pkg/server"
	"entry-notificator/pkg/util"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	config.Init(util.RetrieveEnv("APP_ENV", false))

	database.Init(false)
	defer database.Close()

	server.Init()
}
