package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rifqee23/Mini-Ecommerce-API/internal/domain"
	"github.com/rifqee23/Mini-Ecommerce-API/internal/dto"
	"github.com/rifqee23/Mini-Ecommerce-API/internal/helper"
	"github.com/rifqee23/Mini-Ecommerce-API/internal/repository"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func (service *AuthServiceImpl) RegisterUser(ctx context.Context, request dto.AuthRegisterRequest) (dto.UserResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return dto.UserResponse{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				err = fmt.Errorf("commit failed: %w", commitErr)
			}
		}
	}()

	user := domain.Users{
		Email:    request.Email,
		Password: request.Password,
		Role:     domain.RoleType(request.Role),
	}

	create, err := service.UserRepository.Create(ctx, tx, user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return helper.ToUserResponse(create), nil

}

func (service *AuthServiceImpl) LoginUser(ctx context.Context, request dto.AuthLoginRequest) (dto.AuthLoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (service *AuthServiceImpl) ChangePassword(ctx context.Context, request dto.AuthLoginRequest) (dto.AuthChangePasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (service *AuthServiceImpl) FindByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return dto.UserResponse{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				err = fmt.Errorf("commit failed: %w", commitErr)
			}
		}
	}()

	byEmail, err := service.UserRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return helper.ToUserResponse(byEmail), nil
}
