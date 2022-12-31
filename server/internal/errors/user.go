package errors

import "fmt"

type NoSuchAUserError struct {
	UserID string
}

func (e *NoSuchAUserError) Error() string {
	return fmt.Sprintf("No such a user having the id: %s", e.UserID)
}

type InvalidPasswordError struct {
	UserPassword string
}

func (e *InvalidPasswordError) Error() string {
	return fmt.Sprintf("Invalid password: %s", e.UserPassword)
}
