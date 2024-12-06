package service

import (
	"fibertesttask/internal/model"
	"fibertesttask/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id int) (*model.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.Delete(id)
}
