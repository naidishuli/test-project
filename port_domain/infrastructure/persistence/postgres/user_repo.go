package postgres

import . "port_domain/domain/repositories"

type userRepo struct {
	BaseRepository
}

func NewUserRepo(b BaseRepository) *userRepo {
	return &userRepo{b}
}
