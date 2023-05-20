package utility

import "golang.org/x/crypto/bcrypt"

func HashingFunction(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	
	return string(hashedBytes), nil
}