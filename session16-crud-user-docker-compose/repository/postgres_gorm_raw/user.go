package postgresgormraw

import (
	"context"
	"errors"
	"log"
	"session16-crud-user-docker-compose/entity"

	"gorm.io/gorm"
)

type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	db GormDBIface
}

func NewUserRepository(db GormDBIface) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	query := "INSERT INTO users(name,email,password,created_at,updated_at)VALUES($1,$2,$3,NOW(),NOW()) RETURNING id"
	var createdID int
	if err := r.db.WithContext(ctx).Raw(query, user.Name, user.Email, user.Password).Scan(&createdID).Error; err != nil {
		log.Printf("Error create user : %v\n", err)
		return entity.User{}, err
	}
	user.ID = createdID

	return *user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	query := "Select id,name,email,password,created_at,updated_at From users Where id = $1"

	if err := r.db.WithContext(ctx).Raw(query, id).Scan(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		log.Printf("Error get user : %v\n", err)
		return entity.User{}, nil
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	query := "Update users Set name = $1, email = $2, password = $3, updated_at = NOW() Where id = $4"
	log.Println(query)
	log.Println(id)
	if err := r.db.WithContext(ctx).Exec(query, user.Name, user.Email, user.Password, id).Error; err != nil {
		log.Printf("Error update user : %v\n", err)
		return entity.User{}, nil
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	query := "Delete from users where id = $1"
	log.Println(query)
	log.Println((id))
	if err := r.db.WithContext(ctx).Exec(query, id).Error; err != nil {
		log.Printf("Error deleting user : %v\n", err)
		return err
	}

	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	query := "Select id,name,email,password,created_at,updated_at From users "
	if err := r.db.WithContext(ctx).Raw(query).Scan(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		}
		log.Printf("Error get user : %v\n", err)
		return nil, err
	}

	return users, nil
}
