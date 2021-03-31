package interfaces

import "port_domain/domain/models"

type Services interface {
	AuthService() AuthService
}

type AuthService interface {
	Register(user *models.User) (err error)
	Authenticate()
	RefreshToken()
}
