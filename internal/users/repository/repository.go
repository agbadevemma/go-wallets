package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yourusername/go-wallets/internal/users/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindById(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
}

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) Create(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, full_name, email, password_hash) VALUES (?, ?, ?, ?)", user.Id, user.FullName, user.Email, user.PasswordHash)
	return err
}

func (r *mysqlUserRepository) FindById(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *mysqlUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE email = ?", email).Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mysqlUserRepository) Update(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET full_name = ? WHERE id = ?", user.FullName, user.Id)
	return err
}
