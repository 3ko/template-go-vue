package service

import (
    "mon-projet/internal/domain"
    "mon-projet/internal/repository"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]domain.User, error) {
    return s.repo.FindAll()
}

func (s *UserService) Create(u domain.User) error {
    return s.repo.Create(u)
}
