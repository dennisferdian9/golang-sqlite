package utils

import "os"

func EncryptPassword(password string) string {
	secretKey := os.Getenv("SECRET_KEY")
	return secretKey + password + secretKey
}
