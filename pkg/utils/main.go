package utils

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func init() {
	initEnv()
}

type OTPEmailData struct {
	OTP string
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		panic("Error loading in Env ")
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func ComparePassword(hashedPasswords, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswords), []byte(password))
	return err == nil

}

func GenerateOtp(length int) (string, error) {
	const otpChars = "0123456789"
	otp := make([]byte, length) // Create a byte slice of the specified length.
	_, err := rand.Read(otp)    // Fill the byte slice with random data.
	if err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		otp[i] = otpChars[int(otp[i])%len(otpChars)] // Convert random data to OTP characters.
	}
	return string(otp), nil
}

func loadTemplate(data OTPEmailData) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Construct the full path to the template file
	templatePath := filepath.Join(wd, "pkg", "utils", "templates", "sendOtp.html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func configMail(mail *gomail.Message) {
	smtpHost := os.Getenv("SMTP_HOST")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	d := gomail.NewDialer(smtpHost, 587, username, password)
	if err := d.DialAndSend(mail); err != nil {
		panic(err)
	}
}

func SendOtpMail(email string, otp string) (bool, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("FROM_EMAIL"))
	m.SetHeader("To", email)

	data := OTPEmailData{
		OTP: otp,
	}

	body, err := loadTemplate(data)

	if err != nil {
		log.Fatalf("Could not load email template: %v", err)
	}

	m.SetBody("text/html", body)

	configMail(m)

	log.Println("Email sent successfully!")
	return true, nil
}
