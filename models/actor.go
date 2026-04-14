package models

type Actor struct {
	ID        int
	FirstName string
	LastName  string
	ImageURL  *string //nullabel image may not be available for all actors
}
