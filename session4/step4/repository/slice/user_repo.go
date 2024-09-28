package slice

import "belajargolangpart2/session4/step4/entity"

type IUserRepository interface {
	GetAllUsers() []entity.User
}

type userRepository struct {
	db     []entity.User
	nextID int
}

func NewUserRepository(db []entity.User) IUserRepository {
	return &userRepository{
		db:     db,
		nextID: 1,
	}
}

func (r *userRepository) GetAllUsers() []entity.User {
	return r.db
}
