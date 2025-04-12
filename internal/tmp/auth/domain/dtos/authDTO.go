package dtos

import (
	"errors"
	"strconv"
)

type AuthDTO struct {
	Email      string
	Token      string
	User_id    int
	User_type  string
	User_level string
}

func AuthDTOError(msg string) error {
	return errors.New(msg)
}

func AuthDTOFromArray(items map[string]string) (*AuthDTO, error) {
	keys := []string{"email", "token", "user_id", "user_type", "user_level"}
	for _, key := range keys {
		if _, ok := items[key]; !ok {
			return nil, AuthDTOError("key " + key + " is missing")
		}
	}
	user_id, err := strconv.Atoi(items["user_id"])
	if err != nil {
		return nil, errors.New("user_id should be a valid int")
	}
	return &AuthDTO{
		Email:      items["email"],
		Token:      items["token"],
		User_id:    user_id,
		User_type:  items["user_type"],
		User_level: items["user_level"],
	}, nil
}
