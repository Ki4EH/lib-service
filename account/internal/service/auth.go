package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Ki4EH/lib-service/account/entities"
	"github.com/Ki4EH/lib-service/account/internal/repository"
	"github.com/golang-jwt/jwt"
)

var (
<<<<<<< HEAD
	//salt       = ""
=======
>>>>>>> account
	salt       = os.Getenv("SALT")
	tokenTTL   = 24 * time.Hour
	signingKey = os.Getenv("SIGNINGKEY")
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
<<<<<<< HEAD
	UserId int `json:"user_id"`
=======
	UserId   int `json:"user_id"`
	UserRole int `json:"role"`
>>>>>>> account
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user entities.User) (int, error) {
	user.PasswordHash = generatePasswordHash(user.PasswordHash)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
<<<<<<< HEAD
=======
		user.Flags,
>>>>>>> account
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaisms")
	}
	return claims.UserId, nil
}

func (s *AuthService) ConfirmEmail(token string) error {
	return s.repo.SetActivated(token)
}

func (s *AuthService) ResetPassword(email, newPassword string) error {
	hashNewPassword := generatePasswordHash(newPassword)
	return s.repo.ResetPassword(email, hashNewPassword)
}
