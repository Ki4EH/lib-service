package entities

type User struct {
	Id           int    `json:"-" db:"id"`
	Login        string `json:"login" binding:"required"`
	Email        string `json:"email" binding:"required"`
	PasswordHash string `json:"password_hash" binding:"required"`
}
