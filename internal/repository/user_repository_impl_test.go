package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rifqee23/Mini-Ecommerce-API/internal/domain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestUserRepositoryImpl_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	if err != nil {
		return
	}

	mock.ExpectExec(regexp.QuoteMeta("insert into users (email, password, role) values (?, ?, ?)")).WithArgs("rifqi@gmail.com", "rifqi123", domain.RoleUser).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewUserRepositoryImpl()
	saved, _ := repo.Create(context.Background(), tx, domain.Users{Email: "rifqi@gmail.com", Password: "rifqi123", Role: domain.RoleUser})

	t.Logf("Saved user: %+v", saved)

	assert.Equal(t, 1, saved.UserId)
	assert.Equal(t, "rifqi@gmail.com", saved.Email)
	assert.Equal(t, "rifqi123", saved.Password)
	assert.Equal(t, domain.RoleUser, saved.Role)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryImpl_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	user := domain.Users{
		UserId:   1,
		Email:    "test123@gmail.com",
		Password: "test123",
		Role:     domain.RoleUser,
	}

	mock.ExpectExec(regexp.QuoteMeta("update users set  password=? where id=?")).WithArgs(user.Password, user.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewUserRepositoryImpl()
	updated, _ := repo.Update(context.Background(), tx, user)

	t.Logf("Updated user: %+v", updated)

	assert.Equal(t, "test123", updated.Password)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryImpl_FindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "password", "role", "created_at", "updated_at"}).
		AddRow(1, "rifqi@gmail.com", "rifqi123", domain.RoleUser, time.Now(), time.Now()).
		AddRow(2, "febrianto@gmail.com", "febrianto123", domain.RoleUser, time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta("select user_id, email, password, role, created_at, updated_at from users")).WillReturnRows(rows)

	repo := NewUserRepositoryImpl()
	findAll, _ := repo.FindAll(context.Background(), db)
	t.Logf("data user: %+v", findAll)

	assert.NoError(t, err)
	assert.Equal(t, 1, findAll[0].UserId)
	assert.Equal(t, "rifqi@gmail.com", findAll[0].Email)
	assert.Equal(t, "rifqi123", findAll[0].Password)
	assert.Equal(t, domain.RoleUser, findAll[0].Role)

	assert.Equal(t, 2, findAll[1].UserId)
	assert.Equal(t, "febrianto@gmail.com", findAll[1].Email)
	assert.Equal(t, "febrianto123", findAll[1].Password)
	assert.Equal(t, domain.RoleUser, findAll[1].Role)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryImpl_FindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()

	assert.NoError(t, err)

	rows := sqlmock.NewRows([]string{"user_id", "email", "password", "role", "created_at"}).
		AddRow(1, "rifqi@gmail.com", "rifqi123", domain.RoleUser, time.Now())
	mock.ExpectQuery(regexp.QuoteMeta("select user_id, email, password, role, created_at from users where user_id= ?")).WithArgs(1).WillReturnRows(rows)

	repo := NewUserRepositoryImpl()
	result, _ := repo.FindById(context.Background(), tx, 1)
	t.Logf("data user: %+v", result)

	assert.NoError(t, err)
	assert.Equal(t, 1, result.UserId)
	assert.Equal(t, "rifqi@gmail.com", result.Email)
	assert.Equal(t, "rifqi123", result.Password)
	assert.Equal(t, domain.RoleUser, result.Role)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
