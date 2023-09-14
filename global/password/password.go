package password

import (
	"os"
	"strconv"

	"gin/global/errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	cost, err := strconv.Atoi(os.Getenv("SALT_COST"))

	if err != nil {
		return "", errors.ErrParsingString
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", errors.ErrHashingPassword
	}

	return string(hashedPassword), nil
}

func ComparePassword(passwordReq, passwordDB string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordReq)); err != nil {
		return errors.ErrInvalidPassword
	}
	return nil
}
