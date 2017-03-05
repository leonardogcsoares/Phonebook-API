package repo

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Phonebook TODO
type Phonebook struct {
	gorm.Model
	Firstname string
	Lastname  string
	Nickname  string
	Phones    []Phone
	Emails    []Email
	Addresses []Address
	Socials   []Social
	Birthday  time.Time
	Notes     string
}

// Phone TODO
type Phone struct {
	Number string
	Type   string
}

// Email TODO
type Email struct {
	Value string
	Type  string
}

// Address TODO
type Address struct {
	Value string
	Type  string
}

// Social TODO
type Social struct {
	ID   string
	Type string
}
