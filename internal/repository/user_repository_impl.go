package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rifqee23/Mini-Ecommerce-API/internal/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.Users) (domain.Users, error) {
	query := "insert into users (email, password, role) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Email, user.Password, user.Role)
	if err != nil {
		return domain.Users{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Users{}, err
	}

	user.UserId = int(id)
	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.Users) (domain.Users, error) {
	query := "update users set  password=? where id=?"
	_, err := tx.ExecContext(ctx, query, user.Password, user.UserId)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) ([]domain.Users, error) {
	query := "select user_id, email, password, role, created_at, updated_at from users"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.Users
	for rows.Next() {
		var user domain.Users

		err := rows.Scan(&user.UserId, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Users, error) {
	query := "select user_id, email, password, role, created_at from users where user_id= ?"

	var user domain.Users
	err := tx.QueryRowContext(ctx, query, id).Scan(&user.UserId, &user.Email, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Users{}, errors.New("user not found")
		}
		return domain.Users{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.Users, error) {
	query := "select user_id, email, password, role from users where email=?"

	var user domain.Users
	err := tx.QueryRowContext(ctx, query, email).Scan(&user.UserId, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Users{}, errors.New("user not found")
		}
		return domain.Users{}, err
	}

	return user, nil
}
