package entities

import "fmt"

//User structure
type User struct {
	ID         int64
	Vorname    string
	Abrechnung string
	Status     string
}

//ToString convert struct to a string
func (user User) ToString() string {
	return fmt.Sprintf("id: %d \n name: %s \n Abrechnung: %s \n Status: %s",
		user.ID, user.Vorname, user.Abrechnung, user.Status,
	)
}
