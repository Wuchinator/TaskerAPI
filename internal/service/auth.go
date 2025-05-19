package service

import (
    "errors"
    "github.com/Wuchinator/GoTasker/internal/models"
    "github.com/Wuchinator/GoTasker/internal/repository"
    "golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	//"log"
)

type AuthService struct {
    userRepo *repository.UserRepository
    jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
    return &AuthService{
        userRepo: userRepo,
        jwtSecret: jwtSecret,
    }
}

func (s *AuthService) Register(email, password string) (*models.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Email:        email,
        PasswordHash: string(hashedPassword),
    }

    if err := s.userRepo.CreateUser(user); err != nil {
        return nil, err
    }
    return user, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
    user, err := s.userRepo.GetUserByEmail(email)
    if err != nil {
        return "", err
    }
    if user == nil {
        return "", errors.New("user not found")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })

    return token.SignedString([]byte(s.jwtSecret))
}