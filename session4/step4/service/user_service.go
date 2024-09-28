package service

import (
	"belajargolangpart2/session4/step4/entity"
	"belajargolangpart2/session4/step4/repository/slice"
)

type IUserService interface {
	GetAllUsers() []entity.User
}

type userService struct {
	userRepo slice.IUserRepository
}

func NewUserService(userRepo slice.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() []entity.User {
	return s.userRepo.GetAllUsers()
}
