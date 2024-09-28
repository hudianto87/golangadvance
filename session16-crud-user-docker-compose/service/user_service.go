package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"session16-crud-user-docker-compose/entity"
	postgresgormraw "session16-crud-user-docker-compose/repository/postgres_gorm_raw"
	"time"

	"github.com/redis/go-redis/v9"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

// ngirim data pake tanda *, jika datanya ada yang mau diubah
// ambil data pake tanda &

//untuk menggunakan gorm
// type userService struct {
// 	userRepo postgresgorm.IUserRepository
// }

// untuk menggunakan raw
type userService struct {
	userRepo postgresgormraw.IUserRepository
	rdb      *redis.Client
}

//untuk menggunakan gorm
// func NewUserService(userRepo postgresgorm.IUserRepository) IUserService {
// 	return &userService{userRepo: userRepo}
// }

const redisUserByIDKey = "user:%d"

// untuk menggunakan raw
func NewUserService(userRepo postgresgormraw.IUserRepository, rdb *redis.Client) IUserService {
	return &userService{userRepo: userRepo, rdb: rdb}
}

func (r *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	createdUser, err := r.userRepo.CreateUser(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("error created user: %v", err)
	}

	//serialize user to json
	createdUserJSON, err := json.Marshal(createdUser)

	if err != nil {
		log.Println("failed to mashal user to json", err)
		return createdUser, err
	}

	// set redis key with 1 minute expired
	redisKey := fmt.Sprintf(redisUserByIDKey, createdUser.ID)
	if err := r.rdb.Set(ctx, redisKey, createdUserJSON, 1*time.Minute).Err(); err != nil {
		log.Println("failed to set user to json", err)
		return createdUser, err
	}

	return createdUser, nil
}

func (r *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User

	redisKey := fmt.Sprintf(redisUserByIDKey, id)

	//try to get user to redis cache
	val, err := r.rdb.Get(ctx, redisKey).Result()

	if err == nil {
		// unmarshal the cached
		if err = json.Unmarshal([]byte(val), &user); err == nil {
			return user, nil
		}
		log.Println("failed to unmarshal")
	}

	// if cache miss or unmarshal failed, fetch from DB
	user, err = r.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("error user not found: %v", err)
	}

	//cache the user data with expired
	userJSON, _ := json.Marshal(user)
	if err = r.rdb.Set(ctx, redisKey, userJSON, 1*time.Minute).Err(); err != nil {
		log.Println("failed to set user data to redis", err)
	}

	return user, nil
}

func (r *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	updatedUser, err := r.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("error user not found: %v", err)
	}

	redisKey := fmt.Sprintf(redisUserByIDKey, updatedUser.ID)
	updateJSON, _ := json.Marshal(updatedUser)

	if err = r.rdb.Set(ctx, redisKey, updateJSON, 1*time.Minute).Err(); err != nil {
		log.Println("failed to set user data to redis", err)
	}

	if err := r.rdb.Del(ctx, "all_users").Err(); err != nil {
		log.Println("failed to validate 'all_users'", err)
	}

	return updatedUser, nil
}

func (r *userService) DeleteUser(ctx context.Context, id int) error {
	redisKey := fmt.Sprintf(redisUserByIDKey, id)

	// delete user cache
	if err := r.rdb.Del(ctx, redisKey).Err(); err != nil {
		log.Println("failed to delete cache", err)
	}

	err := r.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("error user not found: %v", err)
	}

	return nil
}

func (r *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	redisKey := "all_users"
	var users []entity.User

	//try to get from redis
	val, err := r.rdb.Get(ctx, redisKey).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &users); err == nil {
			return users, nil
		}
		log.Println("failed to unmarshal")
	}

	//if cache miss or unmarshal
	users, err = r.userRepo.GetAllUsers(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)

	}
	//cache the users
	userJSON, _ := json.Marshal(users)

	if err = r.rdb.Set(ctx, redisKey, userJSON, 1*time.Minute).Err(); err != nil {
		log.Println("failed to set users data to redis")
	}

	if err != nil {
		return nil, fmt.Errorf("error to retrieve data users: %v", err)
	}

	return users, nil
}
