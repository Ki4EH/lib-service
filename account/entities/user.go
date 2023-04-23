package entities

type User struct {
	Id           int    `json:"-" db:"id"`
	Flags        int    `json:"-" db:"flags"`
	Login        string `json:"login" binding:"required" db:"login"`
	Email        string `json:"email" binding:"required" db:"email"`
	PasswordHash string `json:"password_hash" binding:"required" db:"pass_hash"`
}
