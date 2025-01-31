package service

import 	"golang.org/x/crypto/bcrypt"

type PasswordService struct {
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (service *PasswordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (service *PasswordService) ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}