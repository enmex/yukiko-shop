package domain

import "errors"

var (
	UserAlreadyExistsErr = errors.New("user already exists")
	UserNotFoundErr = errors.New("user not found")
	WrongCredentialsErr = errors.New("wrong credentials")
)