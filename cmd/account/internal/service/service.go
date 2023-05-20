package services

import (
	"github.com/Ki4EH/lib-service/account/entities"
	"github.com/Ki4EH/lib-service/account/internal/repository"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
	ConfirmEmail(token string) error
	ResetPassword(email, newPassword string) error
}

type Mail interface {
	SendConfirmMail(user entities.User) (string, error)
	SendPasswordResetMail(email string) (string, error)
}

type Service struct {
	Authorization
	Mail
}

func NewService(repos *repository.Repository, mailConfig MailConfig) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Mail:          NewMailService(mailConfig, repos.Authorization),
	}
}
