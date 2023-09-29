package routes

import (
	"newone/handlers"
	"newone/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.Engine {
	r.POST("/signup", handlers.Signup)
	r.POST("/signupwithotp", handlers.SignupWithOtp)
	r.POST("/signupotpvalidation", handlers.SignupOtpValidation)
	r.POST("/loginwithpassword", handlers.LoginWithPassword)
	r.POST("/loginwithotp", handlers.LoginWithOtp)
	r.POST("/otpvalidation", handlers.LoginOtpValidation)

	r.GET("/home", middlewares.UserRetriveCookie, handlers.Home)  

	return r
}
