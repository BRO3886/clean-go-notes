package pkg

import (
	"errors"
)

//a list of errors
var (
	ErrInvalidToken  = errors.New("Invalid token")
	ErrNotFound      = errors.New("Not found")
	ErrNoContent     = errors.New("No content")
	ErrInvalidSlug   = errors.New("Invalid slug")
	ErrAlreadyExists = errors.New("already exists")
	ErrDatabase      = errors.New("Database error")
	ErrUnauthorised  = errors.New("You are not allowed to perform this action")
	ErrForbidden     = errors.New("Access to this resource is forbidden")
	ErrEmail         = errors.New("Email not valid")
	ErrPassword      = errors.New("Password must be greater than 6 chars")
)
