package service

import (
	"belajargolangpart2/session6dbpgx-crud/entity"
	postgrespgx "belajargolangpart2/session6dbpgx-crud/repository/postgres_pgx"
	"context"
	"fmt"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userService struct {
	userRepo postgrespgx.IUserRepository
}

func NewUserService(userRepo postgrespgx.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (r *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	createdUser, err := r.userRepo.CreateUser(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("error created user: %v", err)
	}

	return createdUser, nil
}

func (r *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	user, err := r.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("error user not found: %v", err)
	}

	return user, nil
}

func (r *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	updatedUser, err := r.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("error user not found: %v", err)
	}

	return updatedUser, nil
}

func (r *userService) DeleteUser(ctx context.Context, id int) error {
	err := r.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("error user not found: %v", err)
	}

	return nil
}

func (r *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	users, err := r.userRepo.GetAllUsers(ctx)

	if err != nil {
		return nil, fmt.Errorf("error to retrieve data users: %v", err)
	}

	return users, nil
}
