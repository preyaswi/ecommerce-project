package repository

import (
	"errors"
	"fmt"
	"newone/entity"
	"newone/infrastructure"
	"newone/models"

	"gorm.io/gorm"
)

func GetByEmail(email string) (*entity.User, error) {
	fmt.Println("🤣")
	var user entity.User
	result := infrastructure.Db.Where(&entity.User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
func GetByPhone(phone string) (*entity.User, error) {
	var user entity.User
	result := infrastructure.Db.Where(&entity.User{Phone: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
func Create(user *entity.User) error {

	return infrastructure.Db.Create(user).Error
}
func CreateSignup(user *models.Signup) error {
	return infrastructure.Db.Create(user).Error

}
func CreateOtpKey(otpKey *entity.OtpKey) error {
	return infrastructure.Db.Create(otpKey).Error
}
func GetByKey(key string) (*entity.OtpKey, error) {
	var otpKey entity.OtpKey
	result:= infrastructure.Db.Where(&entity.OtpKey{Key: key}).First(&otpKey)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &otpKey, nil
}
func GetSignupByPhone(phone string) (*models.Signup, error) {
	var user models.Signup
	result := infrastructure.Db.Where(&models.Signup{Phone: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
func CheckPermission(user *entity.User) (bool, error) {
	result := infrastructure.Db.Where(&entity.User{Phone: user.Phone}).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	permission := user.Permission
	return permission, nil
}
