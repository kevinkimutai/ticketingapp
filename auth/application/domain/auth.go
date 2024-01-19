package domain

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(user User) (User, error) {

	if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		return user, errors.New("missing fields during signup")
	}

	return user, nil
}

func (u User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u User) CreateJWT()        {}
func (u User) VerifyJWT()        {}
func (u User) comparePasswords() {}
