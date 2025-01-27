package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint32   `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Document    string `gorm:"not null"`
	Phone       string `gorm:"not null"`
	DateOfBirth string `gorm:"not null"`
	TenantID    uint32   `gorm:"not null"`
}
