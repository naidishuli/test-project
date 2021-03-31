package app

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"port_domain/app/config"
	_ "port_domain/app/config"
	"port_domain/domain/interfaces"
	"port_domain/domain/models"
	"port_domain/infrastructure/apis"
	"port_domain/infrastructure/logger"
	"port_domain/infrastructure/persistence/postgres"
)

var db *gorm.DB

func init() {
	var err error
	db, err = postgres.NewPostgresConn(postgres.Config{Url: config.Env.DatabaseUrl})
	if err != nil {
		panic(err)
	}
}

type Scope struct {
	interfaces.Repository
	interfaces.Services
	Logger interfaces.Logger
	Client interfaces.HTTPClient
	user   models.User
}

func NewScope() *Scope {
	scp := &Scope{
		Logger:     logger.NewLogger(logrus.Fields{}),
		Repository: NewRepository(db),
	}

	scp.Services = NewServices(scp)
	return scp
}

// SetServices is used for debugging and testing purpose in order to mock
// the available services
func (s *Scope) SetServices(svs interfaces.Services) {
	s.Services = svs
}

func (s *Scope) Log() interfaces.Logger {
	return s.Logger
}

func (s *Scope) HTTPClient() interfaces.HTTPClient {
	if s.Client == nil {
		s.Client = apis.NewHttpClient()
	}

	return s.Client
}
