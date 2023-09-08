package models

type University struct {
	ID       int
	Name     string
	Login    string
	Password string
}

type Faculty struct {
	ID           int
	UniversityID int
	Name         string
}

type Student struct {
	ID   int
	Name string
}
