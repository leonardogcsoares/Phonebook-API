package repo

import (
	"github.com/jinzhu/gorm"
)

// Repo TODO
type Repo interface {
}

// Impl TODO
type Impl struct {
	DB *gorm.DB
}

// NewRepo returns an instance of the repo database methods
func NewRepo(db *gorm.DB) *Impl {
	db.AutoMigrate(&Phonebook{})

	return &Impl{
		DB: db,
	}
}
