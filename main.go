package main

import (
	"github.com/leonardogcsoares/phonebook-api/router"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("db", nil)
	if err != nil {
		return
	}
	defer db.Close()

	r := router.New("3000", db)

	if err := r.Start(); err != nil {
		return
	}
}
