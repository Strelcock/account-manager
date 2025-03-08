package accounts

import "fmt"

const (
	emptyName = iota + 1
	emptyEmail
)

type AccountError struct {
	Name string
	ID   int
}

func (e *AccountError) Error() string {
	return fmt.Sprintf("account error %d: %s", e.ID, e.Name)
}

func (e *AccountError) Unwrap() string {
	return e.Name
}

func NewAccountError(name string, id int) *AccountError {
	return &AccountError{
		Name: name,
		ID:   id,
	}
}
