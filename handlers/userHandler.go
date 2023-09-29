package handlers

import (
	"fmt"
	"net/http"
	"newone/entity"
	"newone/middlewares"
	"newone/models"
	usecase "newone/usecase/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func Signup(c *gin.Context) {
	var userInput models.Signup
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user entity.User
	copier.Copy(&user, &userInput)
	newUser, err := usecase.ExecuteSignup(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)

}
func SignupWithOtp(c *gin.Context) {
	fmt.Println("😊")
	var user models.Signup
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	key, err := usecase.ExecuteSignupWithOtp(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Otp send succesfuly to": user.Phone, "Key": key})
	}
}
func SignupOtpValidation(c *gin.Context) {

	var signupOtpdata models.SignupOtpdata
	if err := c.ShouldBindJSON(&signupOtpdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// key := c.PostForm("key")
	// otp := c.PostForm("otp")
	otp := signupOtpdata.Otp
	key := signupOtpdata.Key

	fmt.Println(key)
	fmt.Println(otp)
	fmt.Println("😢")
	err := usecase.ExecuteSignupOtpValidation(key, otp)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"massage": "user signup succesfull"})
	}

}
func LoginWithOtp(c *gin.Context) {
	type LoginWithOtp struct {
		Phone string `json:"phone"`
	}
	var logindata LoginWithOtp
	if err := c.ShouldBindJSON(&logindata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	phone := logindata.Phone

	fmt.Println(phone)

	// phone := c.PostForm("phone")
	key, err := usecase.ExecuteLogin(phone)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Otp send succesfuly to": phone, "Key": key})
	}

}
func LoginOtpValidation(c *gin.Context) {

	var otpValidation models.OtpValidationData
	if err := c.ShouldBindJSON(&otpValidation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	phone := otpValidation.Phone
	otp := otpValidation.Otp
	key := otpValidation.Key

	fmt.Println(phone)

	user, err := usecase.ExecuteOtpValidation(key, otp)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	middlewares.CreateJwtCookie(user.ID, user.Phone, "user", c)
	c.JSON(http.StatusOK, gin.H{"massage": "user loged in succesfully and cookie stored"})

}

func LoginWithPassword(c *gin.Context) {

	var logindata models.LoginWithPassword
	if err := c.ShouldBindJSON(&logindata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	phone := logindata.Phone
	password := logindata.Password

	fmt.Println(phone)
	fmt.Println(password)

	// var logindata models.LoginWithPassword
	// if err := c.ShouldBindJSON(&logindata); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// phone := logindata.Phone
	// password := logindata.Password

	// fmt.Println(phone)
	// fmt.Println(password)
	userId, err := usecase.ExecuteLoginWithPassword(phone, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		middlewares.CreateJwtCookie(userId, phone, "user", c)
		c.JSON(http.StatusOK, gin.H{"massage": "user loged in succesfully and cookie stored"})
	}

}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"options": "logout - category - products"})
}
