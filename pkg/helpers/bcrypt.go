package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	passwordByte := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func ComparePassword(hashedPassword, password string) bool {
	passwordByte := []byte(password)
	hashedByte := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(hashedByte, passwordByte)
	if err == nil {
		return true
	}

	return false
}
