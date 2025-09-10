package helper

import (
	"github.com/rifqee23/Mini-Ecommerce-API/internal/domain"
	"github.com/rifqee23/Mini-Ecommerce-API/internal/dto"
)

func ToUserResponse(user domain.Users) dto.UserResponse {
	return dto.UserResponse{
		UserId:    user.UserId,
		Email:     user.Email,
		Role:      dto.RoleType(user.Role),
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}
