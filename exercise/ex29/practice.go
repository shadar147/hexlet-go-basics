package ex29

import (
	"encoding/json"
	"errors"
)

type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

var (
	ErrEmailRequired                = errors.New("email is required")                             // когда поле email не заполнено
	ErrPasswordRequired             = errors.New("password is required")                          // когда поле password не заполнено
	ErrPasswordConfirmationRequired = errors.New("password confirmation is required")             // когда поле password_confirmation не заполнено
	ErrPasswordDoesNotMatch         = errors.New("password does not match with the confirmation") // когда поля password и password_confirmation не совпадают
)

func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	r, err := decodeRequest(requestBody)

	if err != nil {
		return CreateUserRequest{}, err
	}

	validateErr := validateRequest(r)

	if validateErr != nil {
		return CreateUserRequest{}, validateErr
	}

	return r, nil
}

func decodeRequest(requestBody []byte) (CreateUserRequest, error) {
	r := CreateUserRequest{}

	err := json.Unmarshal(requestBody, &r)

	return r, err
}

func validateRequest(r CreateUserRequest) error {
	if r.Email == "" {
		return ErrEmailRequired
	}

	if r.Password == "" {
		return ErrPasswordRequired
	}

	if r.PasswordConfirmation == "" {
		return ErrPasswordConfirmationRequired
	}

	if r.Password != r.PasswordConfirmation {
		return ErrPasswordDoesNotMatch
	}

	return nil
}
