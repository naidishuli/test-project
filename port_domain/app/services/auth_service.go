package services

import (
	"port_domain/domain/interfaces"
	"port_domain/domain/models"
)

type authService struct {
	interfaces.AppScope
}

func NewAuthService(scope interfaces.AppScope) *authService {
	return &authService{scope}
}

func (a *authService) Register(user *models.User) (err error) {
	return
}

func (a *authService) Authenticate() {

}

func (a *authService) RefreshToken() {

}
