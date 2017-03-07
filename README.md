# Phonebook API

This project is the solution to one of the challenges presented on a hiring platform which will not be named so as to not encourage cheaters.

## Some Considerations

For this project I've decided on using a couple of Open Source libraries to facilitate development and produce clean and easy to read code. As well as striving for simplicity.
- [Gin](https://github.com/gin-gonic/gin) - for routing and handling of http requests
- [LevelDB](https://github.com/syndtr/goleveldb) - For data persistance I've opted to use a Key/Value store, where data is serialized/deserialized using JSON (Gob could have been used as well, but opted for JSON just cause... =P).

Given the limited time frame + my other responsibilities, this cannot be considered a "professional" or "production ready" solution. It is just demonstration of skill and dominion over Golang.

## Instalation and Usage

Once in the folder at root level. Simply `go run main.go` to start up the server. If necessary `go install` to download the necessary packages.

## Data Modeling

A contact is CRUDed and stored as an `Entry` through the application. It follows the models:

```
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

type Phone struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}

type Email struct {
	Value string `json:"email"`
	Type  string `json:"type"`
}

type Address struct {
	Value string `json:"address"`
	Type  string `json:"type"`
}

type Social struct {
	SID  string `json:"id"`
	Type string `json:"type"`
}
```

Where at least the `Firstname` cannot be empty string and the `Phones` array must contain at least one item.

## Accessing the API

Postman containing possible calls to the API:
`https://www.getpostman.com/collections/5bf2baa67d2dcbc877b0`


#### POST - Login
Necessary to authenticate user and return jwt token:

curl -X POST -H "Content-Type: application/json" -d '{
"username": "admin",
"password": "admin"
}' "http://localhost:3000/login"

#### POST - Create Entry
Given the JWT token returned at login.

This json follows the structure layed out for the `Entry` model up above. Only `Firstname` and one `Phone` entry in the `Phones` array being necessary.

curl -X POST -H "Authorization: jwt token goes here" -H "Content-Type: application/json"  -d '{
	"firstname": "Yasmin",
	"lastname": "Kaufmann",
	"phones": [
		{"number": "07376189706", "type": "mobile"}
	]
}' "http://localhost:3000/phone"

#### GET - Get Entry

curl -X GET -H "Authorization: jwt token goes here" "http://localhost:3000/phone/{id goes here}"

#### PUT - Update Entry

To update an entry, the entire entry must be resubmitted with only the desired fields changed. If an empty body is submitted all fields will then assume their default zero value, except `Firstname` and `Phones` which will maintain their previous state.

curl -X PUT -H "Authorization: jwt goes here" -H "Content-Type: application/json" -d '{
	"firstname": "Yasmin",
	"lastname": "Vilela Kaufmann",
	"phones": [
		{"number": "07356108091", "type": "mobile"}
	]
}' "http://localhost:3000/phone/{id}"

#### DELETE - Delete Entry

curl -X DELETE -H "Authorization: jwt token " -H "Content-Type: application/json" -d '' "http://localhost:3000/phone/{id}"
