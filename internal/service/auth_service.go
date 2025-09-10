package service

import (
	"context"
	"github.com/rifqee23/Mini-Ecommerce-API/internal/dto"
)

type AuthService interface {
	RegisterUser(ctx context.Context, request dto.AuthRegisterRequest) (dto.UserResponse, error)
	LoginUser(ctx context.Context, request dto.AuthLoginRequest) (dto.AuthLoginResponse, error)
	ChangePassword(ctx context.Context, request dto.AuthLoginRequest) (dto.AuthChangePasswordResponse, error)
	FindByEmail(ctx context.Context, email string) (dto.UserResponse, error)
}
