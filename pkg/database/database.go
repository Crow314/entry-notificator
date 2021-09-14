package database

import (
	"entry-notificator/pkg/config"
	"entry-notificator/pkg/util"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init initializes database
func Init() {
	log.Info("Start to initialize database")

	var err error
	db, err = gorm.Open(postgres.Open(retrieveDsn()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database\n%v", err)
	}

	log.Info("Initializing database has completed")
}

// DB returns database connection
func DB() *gorm.DB {
	return db
}

// Close closes database
func Close() {
	db_, err := db.DB()
	if err != nil {
		log.Errorf("Couldn't retrieve database object\n%v", err)
	}

	err = db_.Close()
	if err != nil {
		log.Errorf("Database wasn't closed as expected\n%v", err)
	}
}

func retrieveDsn() (dsn string) {
	conf := config.Config()

	dsn += "host=" + util.RetrieveEnv("POSTGRES_HOST", false) + " "
	dsn += "port=" + util.RetrieveEnv("POSTGRES_PORT", false) + " "
	dsn += "user=" + util.RetrieveEnv("POSTGRES_USER", false) + " "
	dsn += "password=" + util.RetrieveEnv("POSTGRES_PASSWORD", false) + " "
	dsn += "dbname=" + util.RetrieveEnv("POSTGRES_DB", false) + " "
	dsn += "sslmode=" + conf.GetString("database.sslmode") + " "
	dsn += "TimeZone=" + util.RetrieveEnv("TZ", false)

	return
}
