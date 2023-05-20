package repository

import (
	"fmt"

	"github.com/Ki4EH/lib-service/account/entities"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entities.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, email, password_hash, confirm_token) values ($1, $2, $3, $4) RETURNING id", usersTable)
	//todo: email request
	uid := uuid.New().String()
	row := r.db.QueryRow(query, user.Login, user.Email, user.PasswordHash, uid)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}

func (r *AuthPostgres) GetConfirmToken(user entities.User) (string, error) {
	var confirm_token string
	query := fmt.Sprintf("SELECT confirm_token FROM %s WHERE email = $1", usersTable)
	err := r.db.Get(&confirm_token, query, user.Email)
	return confirm_token, err
}

func (r *AuthPostgres) SetActivated(token string) error {
	query := fmt.Sprintf("UPDATE %s SET confirm_token = NULL WHERE confirm_token = '%s'", usersTable, token)
	r.db.QueryRow(query)
	return nil
}

// todo
func (r *AuthPostgres) ResetPassword(email, newPassword string) error {
	return nil
}
