package models

import "strings"

const (
	// ErrNotFound is returned when a resource cannot be found in
	// the database.
	ErrNotFound modelError = "models: resource not found"

	// ErrIDInvalid is returned when an invalid ID is provided.
	ErrIDInvalid modelError = "models: ID provided is invalid"

	// ErrPasswordIncorrect is returned when the password is invalid.
	ErrPasswordIncorrect modelError = "models: incorrect password provided"

	// ErrEmailRequired is returned when an email address is not provide
	// when creating a user
	ErrEmailRequired modelError = "models: email address is required"

	// ErrEmailInvalid is returned when an email address
	// does not match any of our requirements
	ErrEmailInvalid modelError = "models: email address is invalid"

	// ErrEmailTaken is returned when an update() or create() is
	// attempted with an existing email address.
	ErrEmailTaken modelError = "models: email is already taken"

	// ErrPasswordTooShort is returned when the password length
	// does not match the required minimum.
	ErrPasswordTooShort modelError = "models: password must be at least 8 char long"

	// ErrPasswordRequired is returned when a password is not provided.
	ErrPasswordRequired modelError = "models: password is required"

	// ErrRememberTooShort is returned if the remember token is not
	// at least 32 bytes
	ErrRememberTooShort modelError = "models: remember token must be 32 bytes or longer"

	// ErrRememberHashRequired is returned when a create or update is
	// attempter without a user remember token hash
	ErrRememberHashRequired modelError = "models: remember required"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	return strings.Title(s)
}
