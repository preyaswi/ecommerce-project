package models

type Signup struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}
type SignupOtpdata struct {
	Key string `json:"key"`
	Otp string `json:"otp"`
}
type LoginWithPassword struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type OtpValidationData struct {
	Key    string `json:"key"`
	Otp    string `json:"otp"`
	Phone  string `json:"phone"`
}
