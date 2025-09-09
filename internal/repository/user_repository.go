package repository

import (
	"context"
	"database/sql"
	"github.com/rifqee23/Mini-Ecommerce-API/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.Users) (domain.Users, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.Users) (domain.Users, error)
	FindAll(ctx context.Context, db *sql.DB) ([]domain.Users, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Users, error)
}
