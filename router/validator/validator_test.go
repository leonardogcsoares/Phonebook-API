package validator

import (
	"testing"

	"github.com/leonardogcsoares/phonebook-api/router/repo"
)

func TestIsValidEntry(t *testing.T) {
	testCases := []struct {
		Name     string
		Entry    repo.Entry
		Expected error
	}{
		{
			Name:     "empty entry",
			Entry:    repo.Entry{},
			Expected: ErrFirstnameEmpty,
		},
		{
			Name: "no phones submitted",
			Entry: repo.Entry{
				Firstname: "Leonardo",
			},
			Expected: ErrPhonesEmpty,
		},
		{
			Name: "succesful entry",
			Entry: repo.Entry{
				Firstname: "Leonardo",
				Phones: []repo.Phone{
					repo.Phone{
						Number: "031997198788",
						Type:   "mobile",
					},
				},
			},
			Expected: nil,
		},
	}

	i := Impl{}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if err := i.IsValidEntry(tc.Entry); err != tc.Expected {
				t.Errorf("got '%s' want '%s'", err.Error(), tc.Expected.Error())
			}
		})
	}
}
