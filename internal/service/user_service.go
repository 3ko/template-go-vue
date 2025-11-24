package service

import (
	"mon-projet/internal/db"
	"mon-projet/internal/model"
)

type UserService struct {
	repo *db.UserRepository
}

func NewUserService(repo *db.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) Create(u model.User) error {
	return s.repo.Create(u)
}

func (s *UserService) GetByID(id int64) (model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) Update(id int64, u model.User) error {
	return s.repo.Update(id, u)
}

func (s *UserService) Delete(id int64) error {
	return s.repo.Delete(id)
}
