package interfaces

import "client_api/domain/models"

type Services interface {
	AuthService() AuthService
}

type AuthService interface {
	Register(user *models.User) (err error)
	Authenticate()
	RefreshToken()
}
