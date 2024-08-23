package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Import the bcrypt package at the top of your file
	// import "golang.org/x/crypto/bcrypt"

	// Generate a hashed version of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CompPassword(hashPassword string, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))

	return err == nil
}
