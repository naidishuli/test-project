package interfaces

import "port_domain/domain/repositories"

type Repository interface {
	UserRepository() repositories.UserRepository
}
