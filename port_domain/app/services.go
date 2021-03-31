package app

import (
	"port_domain/app/services"
	"port_domain/domain/interfaces"
)

type Services struct {
	scope interfaces.AppScope
	Auth  interfaces.AuthService
}

func NewServices(sc interfaces.AppScope) *Services {
	return &Services{
		scope: sc,
	}
}

func (s *Services) AuthService() interfaces.AuthService {
	if s.Auth == nil {
		s.Auth = services.NewAuthService(s.scope)
	}

	return s.Auth
}
