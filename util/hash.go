package util

import (
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func VerifyHash(unhashed string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(unhashed))

	return err == nil
}
