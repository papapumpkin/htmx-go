package services

import (
	"htmx-go/internals/models"
	"htmx-go/internals/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	existingUser, err := s.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email

	return s.UserRepo.UpdateUser(existingUser)
}

func (s *UserService) DeleteUser(id uint) error {
	_, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	return s.UserRepo.DeleteUser(id)
}
