package utils

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	TWILIO_ACCOUNT_SID string
	TWILIO_AUTH_TOKEN  string
	VERIFY_SERVICE_SID string
	client             *twilio.RestClient
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	TWILIO_ACCOUNT_SID = os.Getenv("KEY1")
	TWILIO_AUTH_TOKEN = os.Getenv("KEY2")
	VERIFY_SERVICE_SID = os.Getenv("KEY3")
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})
}

func SendOtp(phone string) (string, error) {
	to := "+91" + phone
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("Otp failed to generate")
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
		return *resp.Sid, nil
	}
}

func CheckOtp(phone, code string) error {

	to := "+91" + phone
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)
	fmt.Println("🙌")
	fmt.Println(to)
	fmt.Println("🙌")

	fmt.Println(code)
	fmt.Println("🙌")

	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)
	fmt.Println(resp, "👌")

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Invalid otpppppp")
	} else if *resp.Status == "approved" {
		return nil
	} else {
		return errors.New("Invalid otpp")
	}

	// if err != nil {
	// 	fmt.Println("Twilio API error:", err.Error())
	// 	return errors.New("Error verifying OTP")
	// } else if resp.Status == nil || *resp.Status != "approved" {
	// 	fmt.Println("OTP verification failed")
	// 	return errors.New("Invalid OTP")
	// }
}
