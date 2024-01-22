package db

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kevinkimutai/ticketingapp/auth/application/domain"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dbString string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&User{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) CreateUser(user domain.User) (domain.User, error) {
	err := a.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (a *Adapter) Login(user domain.LoginUser) (string, error) {

	//Check Email
	foundUser := User{}
	err := a.db.Where("email = ?", user.Email).First(&foundUser).Error
	if err != nil {
		return "", status.Errorf(codes.PermissionDenied, "wrong email or password")
	}

	//Compare Passwords
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return "", status.Errorf(codes.PermissionDenied, "wrong email or password")
	}

	//Create JWT
	token, err := foundUser.CreateJWT()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (user *User) CreateJWT() (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		errMsg := fmt.Sprintf("something went wrong when hashing the token: %v", err)
		return "", status.Errorf(codes.Internal, errMsg)
	}

	return tokenString, nil

}
