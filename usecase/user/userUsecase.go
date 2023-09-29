package usecase

import (
	"errors"
	"fmt"
	"newone/entity"
	"newone/models"
	repository "newone/repository/user"
	"newone/utils"

	"golang.org/x/crypto/bcrypt"
)

func ExecuteSignup(user entity.User) (*entity.User, error) {
	fmt.Println(user)
	email, err := repository.GetByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return nil, errors.New("error with server")
	}
	if email != nil {
		return nil, errors.New("user with this email already exists")
	}
	phone, err := repository.GetByPhone(user.Phone)
	if err != nil {
		return nil, errors.New("error with server")
	}
	if phone != nil {
		return nil, errors.New("user with this phone no already exists")
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  string(hasedPassword),
	}
	err = repository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
func ExecuteSignupWithOtp(user models.Signup) (string, error) {
	fmt.Println("😂")
	var otpKey entity.OtpKey
	email, err := repository.GetByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return "", errors.New("error with server")
	}
	if email != nil {
		return "", errors.New("user with this email already existsssss")
	}
	phone, err := repository.GetByPhone(user.Phone)
	if err != nil {
		return "", errors.New("error with server")
	}
	if phone != nil {
		return "", errors.New("user with this phone no already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)
	fmt.Println(user)
	key, err := utils.SendOtp(user.Phone)
	if err != nil {
		return "", err
	} else {

		err = repository.CreateSignup(&user)
		otpKey.Key = key
		otpKey.Phone = user.Phone

		err = repository.CreateOtpKey(&otpKey)
		if err != nil {
			return "", err
		}
		return key, nil
	}
}
func ExecuteSignupOtpValidation(key string, otp string) error {
	fmt.Println(otp,"😊")

	result, err := repository.GetByKey(key)
	if err != nil {
		return err
	}
	user, err := repository.GetSignupByPhone(result.Phone)
	if err != nil {
		return err
	}
	err = utils.CheckOtp(result.Phone, otp)
	if err != nil {
		return err
	} else {
		newUser := &entity.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Password:  user.Password,
		}

		err1 := repository.Create(newUser)
		if err1 != nil {
			return err1
		} else {
			return nil
		}
	}

}
func ExecuteLogin(phone string) (string, error) {
	var otpKey entity.OtpKey
	result, err := repository.GetByPhone(phone)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", errors.New("user with this phone not found")
	}
	permission, err := repository.CheckPermission(result)
	if permission == false {
		return "", errors.New("user permission denied")
	}
	key, err := utils.SendOtp(phone)
	if err != nil {
		return "", err
	} else {
		otpKey.Key = key
		otpKey.Phone = phone
		err = repository.CreateOtpKey(&otpKey)
		if err != nil {
			return "", err
		}
		return key, nil
	}

}
func ExecuteOtpValidation(key, otp string) (*entity.User, error) {
	result, err := repository.GetByKey(key)
	if err != nil {
		return nil, err
	}
	user, err := repository.GetByPhone(result.Phone)
	if err != nil {
		return nil, err
	}
	err1 := utils.CheckOtp(result.Phone, otp)
	if err1 != nil {
		return nil, err1
	}
	return user, nil
}
func ExecuteLoginWithPassword(phone, password string) (int, error) {
	user, err := repository.GetByPhone(phone)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errors.New("user with this phone not found")
	}
	permission, err := repository.CheckPermission(user)
	if permission == false {
		return 0, errors.New("user permission denied")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return 0, errors.New("invalid Password")
	} else {
		return user.ID, nil
	}

}
