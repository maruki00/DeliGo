package valueobjects

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func NewPassword(password string) (Password, error) {
	if len(password) < 8 {
		return "", errors.New("password should contains atleast 88 chars")
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[^A-Za-z0-9]`).MatchString(password)
	if !(hasLower && hasUpper && hasDigit && hasSpecial) {
		return "", errors.New("password should contains upper cases and lower cases and number and symboles")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return Password(string(hashed)), nil
}

func (_this Password) Verify(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(_this), []byte(pwd))
	return err == nil
}

func (_this *Password) String() string {
	return string(*_this)
}
