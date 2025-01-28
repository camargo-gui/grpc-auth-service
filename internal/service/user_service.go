package service

import (
	"grpc-auth-service/internal/generated/auth"
	"grpc-auth-service/internal/model"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (service *UserService) CreateUser(req *auth.RegisterRequest) (uint32, error) {
	user := model.User{
		Name:        req.GetName(),
		Email:       req.GetEmail(),
		Password:    req.GetPassword(),
		Document:    req.GetDocument(),
		Phone:       req.GetPhone(),
		DateOfBirth: req.GetDateOfBirth(),
		TenantID:    req.GetTenantId(),
	}

	result := service.DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}
