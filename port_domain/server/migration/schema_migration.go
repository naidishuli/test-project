package main

import (
	"github.com/sirupsen/logrus"
	"port_domain/app/config"
	_ "port_domain/app/config"
	"port_domain/domain/models"
	"port_domain/infrastructure/logger"
	"port_domain/infrastructure/persistence/postgres"
)

func init() {

}

func main() {
	l := logger.NewLogger(logrus.Fields{})
	db, err := postgres.NewPostgresConn(postgres.Config{Url: config.Env.DatabaseUrl})
	if err != nil {
		l.Panic(err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Port{},
	)

	if err != nil {
		l.Panic(err)
	}

	if db.Error != nil {
		l.Panic(err)
	}
}
