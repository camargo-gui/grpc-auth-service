package service

import (
	"grpc-auth-service/internal/generated/auth"
	"grpc-auth-service/internal/model"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
	PasswordService *PasswordService
}

func NewUserService(db *gorm.DB, passwordService *PasswordService) *UserService {
	return &UserService{
		DB: db,
		PasswordService: passwordService,
	}
}

func (service *UserService) CreateUser(req *auth.RegisterRequest) (uint32, error) {

	password, err := service.PasswordService.HashPassword(req.GetPassword())
	if err != nil {
		return 0, err
	}

	user := model.User{
		Name:        req.GetName(),
		Email:       req.GetEmail(),
		Password:    string(password),
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

func (service *UserService) GetUserByEmail(email string, tenant uint32) (*model.User, error) {
	var user model.User
	result := service.DB.Where("email = ? AND tenant_id = ?", email, tenant).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}