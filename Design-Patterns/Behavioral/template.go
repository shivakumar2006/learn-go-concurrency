package main

import (
	"fmt"
)

type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOtp(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

// SMS
type Sms struct {
	Otp
}

func (s *Sms) genRandomOTP(len int) string {
	randomOtp := "1234"
	fmt.Printf("SMS: generating random otp: %s\n", randomOtp)
	return randomOtp
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

// email
type Email struct {
	Otp
}

func (e *Email) genRandomOTP(len int) string {
	randomOtp := "1234"
	fmt.Printf("EMAIL: generating random otp : %s\n", randomOtp)
	return randomOtp
}

func (e *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp : %s to cache \n", otp)
}

func (e *Email) getMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (e *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email : %s\n", message)
	return nil
}

func main() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOtp(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOtp(4)
}
