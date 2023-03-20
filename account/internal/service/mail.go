package services

import (
	"fmt"
	"strconv"

	"github.com/Ki4EH/lib-service/account/entities"
	"github.com/Ki4EH/lib-service/account/internal/repository"
	"github.com/go-gomail/gomail"
	"github.com/spf13/viper"
)

type MailConfig struct {
	Email      string
	SMTPPort   string
	SMTPServer string
	Password   string
}

type MailService struct {
	Config MailConfig
	Repos  repository.Authorization
}

func NewMailService(cfg MailConfig, repos repository.Authorization) *MailService {
	return &MailService{Config: cfg, Repos: repos}
}

func (m *MailService) SendConfirmMail(user entities.User) (string, error) {
	confirmToken, err := m.Repos.GetConfirmToken(user)
	if err != nil {
		return "", nil
	}
	msg := (fmt.Sprintf("<h1>lib-service подтверждение</h1><div>Перейдите по <a href=\"%s\"> ссылке </a> для подтверждении почты </div>\r\n", viper.GetString("SERVER_URL")+"/api/verify/"+confirmToken))
	ms := gomail.NewMessage()
	ms.SetHeader("From", m.Config.Email)
	ms.SetHeader("To", user.Email)
	ms.SetHeader("Subject", "Подтверждение почты lib-service!")
	ms.SetBody("text/html", msg)
	port, err := strconv.Atoi(m.Config.SMTPPort)
	if err != nil {
		return "", err
	}
	d := gomail.NewDialer(m.Config.SMTPServer, port, m.Config.Email, m.Config.Password)
	if err := d.DialAndSend(ms); err != nil {
		return "", err
	}
	return confirmToken, nil
}

// todo
func (m *MailService) SendPasswordResetMail(email string) (string, error) {
	return "newPassword", nil
}
