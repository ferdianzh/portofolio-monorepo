package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(text string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(hash)
}

func ComparePassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}