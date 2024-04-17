package service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"hrm/internal/domain"
	"hrm/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *domain.CreateUserRequest) error {
	_, err := s.userRepository.GetUserByUserName(ctx, req.UserName)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	hashPass, err := HashPass(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashPass

	return s.userRepository.CreateUser(ctx, &repository.User{
		UserName: req.UserName,
		Password: req.Password,
		FullName: req.Name,
		Phone:    req.Phone,
		Email:    req.Email,
	})
}

func (s *UserService) GetUser(ctx context.Context, req *domain.GetUserRequest) ([]*repository.User, error) {
	users, err := s.userRepository.GetUser(ctx, &repository.GetUserRequest{
		Username: req.UserName,
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *domain.UpdateUserRequest) error {
	hashPass, err := HashPass(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashPass

	return s.userRepository.UpdateUser(ctx, &repository.User{
		UserName: req.UserName,
		Password: req.Password,
		FullName: req.Name,
		Phone:    req.Phone,
		Email:    req.Email,
	})

}
