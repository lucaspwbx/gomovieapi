package models

type Movie struct {
	Title  string
	Year   int
	Actors []*Actor
}

type Actor struct {
	Name string
	Age  int
}
