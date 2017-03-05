package main

import "github.com/leonardogcsoares/phonebook-api/router"

func init() {

}

func main() {

	r := router.New("3000")

	if err := r.Start(); err != nil {
		return
	}
}
