package app

import (
	_ "client_api/app/config"
	"client_api/domain/interfaces"
	"client_api/domain/models"
	"client_api/infrastructure/apis"
	"client_api/infrastructure/logger"
	"github.com/sirupsen/logrus"
)

func init() {

}

type Scope struct {
	interfaces.Services
	Logger interfaces.Logger
	Client interfaces.HTTPClient
	user   models.User
}

func NewScope() *Scope {
	scp := &Scope{
		Logger:     logger.NewLogger(logrus.Fields{}),
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