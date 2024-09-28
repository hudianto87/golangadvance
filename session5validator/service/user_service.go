package service

import (
	"belajargolangpart2/session5validator/entity"
	"belajargolangpart2/session5validator/repository/slice"
	"fmt"
)

type IUserService interface {
	CreateUser(user *entity.User) entity.User
	GetUserByID(id int) (entity.User, error)
	UpdateUser(id int, user entity.User) (entity.User, error)
	DeleteUser(id int) error
	GetAllUsers() []entity.User
}

type userService struct {
	userRepo slice.IUserRepository
}

func NewUserService(userRepo slice.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (r *userService) CreateUser(user *entity.User) entity.User {
	return r.userRepo.CreateUser(user)
}

func (r *userService) GetUserByID(id int) (entity.User, error) {
	user, found := r.userRepo.GetUserByID(id)
	if !found {
		return entity.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *userService) UpdateUser(id int, user entity.User) (entity.User, error) {
	updatedUser, found := r.userRepo.UpdateUser(id, user)
	if !found {
		return entity.User{}, fmt.Errorf("user not found")
	}

	return updatedUser, nil
}

func (r *userService) DeleteUser(id int) error {
	if !r.userRepo.DeleteUser(id) {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (s *userService) GetAllUsers() []entity.User {
	return s.userRepo.GetAllUsers()
}
