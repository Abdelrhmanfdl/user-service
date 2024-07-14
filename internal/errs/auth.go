package errs

import "fmt"

type UserExisting struct {
	Message string
}

func (e *UserExisting) Error() string {
	return fmt.Sprintf("failed to create account: %s", e.Message)
}

type HashingError struct {
	Message string
}

func (e *HashingError) Error() string {
	return fmt.Sprintf("failed to create account: %s", e.Message)
}

type WrongEmailOrPassword struct {
	Message string
}

func (e *WrongEmailOrPassword) Error() string {
	return fmt.Sprintf("failed to login: %s", e.Message)
}
