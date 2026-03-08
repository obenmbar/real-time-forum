package forum

import (
	"fmt"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error en genarte password")
	}
	return string(hashpassword), nil
}

func GenerateUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("error en generate uuid")
	}
	return id.String(), nil
}
