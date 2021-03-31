package app

import (
	"gorm.io/gorm"
	. "port_domain/domain/repositories"
	"port_domain/infrastructure/persistence/postgres"
)

type Repository struct {
	Base BaseRepository
	User UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Base: postgres.NewBaseRepo(db),
	}
}

func (r *Repository) UserRepository() UserRepository {
	if r.User == nil {
		r.User = postgres.NewUserRepo(r.Base.New())
	}

	return r.User
}
