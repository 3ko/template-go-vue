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

func (s *UserService) GetByID(id int64) (domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) Update(id int64, u domain.User) error {
	return s.repo.Update(id, u)
}

func (s *UserService) Delete(id int64) error {
	return s.repo.Delete(id)
}
