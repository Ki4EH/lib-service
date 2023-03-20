package repository

import (
	"github.com/Ki4EH/lib-service/account/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(username, password string) (entities.User, error)
	SetActivated(token string) error
	GetConfirmToken(user entities.User) (string, error)
	ResetPassword(email, password string) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
