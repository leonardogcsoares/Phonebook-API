package repo

// Entry TODO
type Entry struct {
	ID        string    `json:"id"`
	Firstname string    `gorm:"not null" json:"firstname"`
	Lastname  string    `json:"lastname"`
	Nickname  string    `json:"nickname"`
	Phones    []Phone   `gorm:"not null" json:"phones"`
	Emails    []Email   `json:"emails"`
	Addresses []Address `json:"addresses"`
	Socials   []Social  `json:"socials"`
	Notes     string    `json:"notes"`
}

// Phone TODO
type Phone struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}

// Email TODO
type Email struct {
	Value string `json:"email"`
	Type  string `json:"type"`
}

// Address TODO
type Address struct {
	Value string `json:"address"`
	Type  string `json:"type"`
}

// Social TODO
type Social struct {
	SID  string `json:"id"`
	Type string `json:"type"`
}
