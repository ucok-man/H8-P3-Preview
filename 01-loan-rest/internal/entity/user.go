package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   int `gorm:"primaryKey"`
	Username string
	Email    string
	Password []byte
}

func (p *User) SetPassword(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.Password = hash
	return nil
}

func (p *User) MatchesPassword(plaintextPassword string) error {
	err := bcrypt.CompareHashAndPassword(p.Password, []byte(plaintextPassword))
	if err != nil {
		return err
	}

	return nil
}
