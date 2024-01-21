package domain

import (
	"errors"
	"net/mail"
	"regexp"
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

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(user User) (User, error) {

	if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		return user, errors.New("missing fields during signup")
	}

	return user, nil
}

func NewLogin(user LoginUser) (LoginUser,error){
	if user.Email == "" || user.Password == ""  {
		return user, errors.New("missing fields during Login")
	}

	return user, nil

}

func (u User) HashPassword() (User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return u, err
	}

	return User{FirstName: u.FirstName, LastName: u.LastName, Email: u.Email, Password: string(bytes), CreatedAt: u.CreatedAt}, nil
}

func CheckEmail(email string) error {
	_, err := mail.ParseAddress(email)

	return err

}

func (u User) CheckPasswordStrength() error {
	// Check if the password is at least 8 characters long
	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Check if the password contains at least one uppercase letter
	if ok, _ := regexp.MatchString("[A-Z]", u.Password); !ok {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check if the password contains at least one lowercase letter
	if ok, _ := regexp.MatchString("[a-z]", u.Password); !ok {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Check if the password contains at least one digit
	if ok, _ := regexp.MatchString("[0-9]", u.Password); !ok {
		return errors.New("password must contain at least one digit")
	}

	// Check if the password contains at least one special character
	if ok, _ := regexp.MatchString("[!@#$%^&*(),.?\":{}|<>]", u.Password); !ok {
		return errors.New("password must contain at least one special character")
	}

	// Password meets all criteria
	return nil
}

func (u User) CreateJWT()        {}
func (u User) VerifyJWT()        {}
func (u User) comparePasswords() {}
