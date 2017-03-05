package repo

import (
	"encoding/json"

	"github.com/syndtr/goleveldb/leveldb"
)

// Repo TODO
type Repo interface {
	CreateEntry(Entry) (Entry, error)
	GetEntry(string) (Entry, error)
	UpdateEntry(string, Entry) (Entry, error)
	DeleteEntry(string) error
}

// Impl TODO
type Impl struct {
	DB *leveldb.DB
}

// NewRepo returns an instance of the repo database methods
func NewRepo(db *leveldb.DB) *Impl {
	return &Impl{
		DB: db,
	}
}

// CreateEntry TODO
func (i Impl) CreateEntry(entry Entry) (Entry, error) {
	payload, err := json.Marshal(entry)
	if err != nil {
		return entry, err
	}

	err = i.DB.Put([]byte(entry.ID), payload, nil)
	if err != nil {
		return entry, err
	}

	return entry, nil
}

// GetEntry TODO
func (i Impl) GetEntry(id string) (Entry, error) {

	var e Entry
	resp, err := i.DB.Get([]byte(id), nil)
	if err != nil {
		return e, err
	}

	err = json.Unmarshal(resp, &e)
	if err != nil {
		return e, err
	}

	return e, nil
}

// UpdateEntry TODO
func (i Impl) UpdateEntry(id string, entry Entry) (Entry, error) {
	var e Entry
	resp, err := i.DB.Get([]byte(id), nil)
	if err != nil {
		return e, err
	}

	err = json.Unmarshal(resp, &e)
	if err != nil {
		return e, err
	}

	prevID := e.ID
	previousName := e.Firstname
	previousPhones := e.Phones

	e = entry
	e.ID = prevID
	if e.Firstname == "" {
		e.Firstname = previousName
	}
	if len(e.Phones) == 0 {
		e.Phones = previousPhones
	}

	payload, err := json.Marshal(e)
	if err != nil {
		return entry, err
	}

	err = i.DB.Put([]byte(e.ID), payload, nil)
	if err != nil {
		return entry, err
	}

	return e, nil
}

// DeleteEntry TODO
func (i Impl) DeleteEntry(id string) error {

	err := i.DB.Delete([]byte(id), nil)
	if err != nil {
		return err
	}

	return nil
}
