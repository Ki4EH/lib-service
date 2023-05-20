package main

import (
	"log"
	"os"

	"github.com/Ki4EH/lib-service/account/internal/handler"
	"github.com/Ki4EH/lib-service/account/internal/repository"
	"github.com/Ki4EH/lib-service/account/internal/server"
	services "github.com/Ki4EH/lib-service/account/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := gotenv.Load("../../.env"); err != nil {
		logrus.Fatal(err)
	}
	if err := initConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	postgres, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USERNAME"),
		Password: os.Getenv("DB_PASS"),
		DBName:   viper.GetString("DB_NAME"),
		SSLMode:  viper.GetString("DB_SSLMODE"),
	})
	smtp := services.MailConfig{
		Email:      viper.GetString("EMAIL"),
		SMTPPort:   viper.GetString("SMTP_PORT"),
		SMTPServer: viper.GetString("SMTP_SERVER"),
		Password:   os.Getenv("MAIL_PASS"),
	}

	if err != nil {
		log.Fatal(err)
	}
	db := repository.NewRepository(postgres)
	if err != nil {
		log.Fatal(err)
	}
	service := services.NewService(db, smtp)
	if err != nil {
		log.Fatal(err)
	}
	srv := new(server.Server)
	handlers := handler.New(service)
	if err := srv.Run(viper.GetString("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}

}

func initConfig() error {
	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
