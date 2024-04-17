package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"hrm/internal/domain"
	"hrm/repository"
)

type AuthService struct {
	authRepo *repository.AuthRepository
	userRepo *repository.UserRepository
}

func NewAuthService(authRepo *repository.AuthRepository, userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error) {
	userExisted, err := s.userRepo.GetUserByUserName(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if !CheckPasswordHash(req.Password, userExisted.Password) {
		return nil, fmt.Errorf("password is incorrect")
	}

	accessToken, expireTime, err := createSession(req.Username)
	if err != nil {
		return nil, err
	}

	err = s.authRepo.CreateSession(ctx, &repository.Session{
		UserName: req.Username,
		Expired:  expireTime,
		Token:    accessToken,
		IsLogout: nil,
	})

	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token: accessToken,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	ar := strings.SplitAfter(token, "Bearer ")
	if len(ar) != 2 {
		return fmt.Errorf("token inValid")
	}
	token = strings.TrimSpace(ar[1])
	userInfo, err := parseToken(token, []byte("secret"))
	if err != nil {
		return err
	}

	err = s.authRepo.Logout(ctx, userInfo.UserName)
	if err != nil {
		return err
	}
	return nil
}

func parseToken(token string, secretKey []byte) (*domain.Claims, error) {
	accessToken, err := jwt.Parse(strings.Trim(token, " "), func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if accessToken.Claims.Valid() != nil {
		return nil, fmt.Errorf("can't parse token")
	}

	var claim domain.Claims
	tokenString, err := json.Marshal(accessToken.Claims)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tokenString, &claim)
	if err != nil {
		return nil, err
	}
	t := time.Now()
	if claim.ExpiresAt < t.Unix() {
		return nil, fmt.Errorf("token is expire")
	}

	return &claim, nil
}

func HashPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
