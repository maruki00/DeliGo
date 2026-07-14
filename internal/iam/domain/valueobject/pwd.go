package valueobject

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type PWD string

func NewPWD(password string) (PWD, error) {
	if len(password) < 8 {
		return "", errors.New("password should contains atleast 8 chars")
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
	return PWD(string(hashed)), nil
}

func (_this PWD) Verify(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(_this), []byte(pwd))
	return err == nil
}

func (_this *PWD) String() string {
	return string(*_this)
}
