package validator

import (
	"fmt"

	"github.com/leonardogcsoares/phonebook-api/router/repo"
)

var (
	// ErrInvalidID TODO
	ErrInvalidID = fmt.Errorf("invalid id provided")
	// ErrFirstnameEmpty TODO
	ErrFirstnameEmpty = fmt.Errorf("first name cannot be empty string")
	// ErrPhonesEmpty TODO
	ErrPhonesEmpty = fmt.Errorf("phones cannot be empty array")
)

// V TODO
type V interface {
	IsValidEntry(repo.Entry) error
	IsValidID(string) error
}

// Impl TODO
type Impl struct {
}

// IsValidEntry TODO
func (i Impl) IsValidEntry(entry repo.Entry) error {

	if entry.Firstname == "" {
		return ErrFirstnameEmpty
	}

	if len(entry.Phones) == 0 {
		return ErrPhonesEmpty
	}

	return nil
}

// IsValidID TODO
func (i Impl) IsValidID(id string) error {
	return nil
}
