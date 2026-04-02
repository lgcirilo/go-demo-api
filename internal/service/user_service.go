package service

import (
	"github.com/lgcirilo/go-demo-api/internal/models"
	"github.com/lgcirilo/go-demo-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}
