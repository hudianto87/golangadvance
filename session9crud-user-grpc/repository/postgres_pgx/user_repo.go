package postgrespgx

import (
	"belajargolangpart2/session9crud-user-grpc/entity"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PgxPoolIface interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	db PgxPoolIface
}

func NewUserRepository(db PgxPoolIface) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	query := "INSERT into users(name,email,password,created_at,updated_at) values($1,$2,$3, NOW(),NOW()) RETURNING id"
	var id int
	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		log.Printf("error creating user: %v\n", err)
		return entity.User{}, err
	}
	user.ID = id
	return *user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {

	var user entity.User
	query := "select id, name, created_at, updated_at from users where id = $1"
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	query := "Update users set name = $1, email = $2, updated_at = NOW() where id = $3"
	_, err := r.db.Exec(ctx, query, user.Name, user.Email, id)
	if err != nil {
		log.Printf("error updating user: %v\n", err)
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	query := "Delete users where id = $1"
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		log.Printf("error deleting user: %v\n", err)
		return err
	}

	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	query := "select id,name,email,created_at, updated_at from users"
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		log.Printf("error getting all users: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	// * itu pointer insert ke memory, & itu pointer mengambil dari memory
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("error scanning users: %v\n", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
