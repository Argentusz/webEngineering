package models

type University struct {
	ID       int
	Name     string
	Login    string
	Password string
	Lang     string
}

type Faculty struct {
	ID           int
	UniversityID int
	Name         string
	ExamDate     string
	ExamAud      string
}

type Student struct {
	ID   int
	Name string
	Exam int
}

type FacultyToStudent struct {
	Fid int
	Sid int
}

type Settings struct {
	Uid  int
	Lang string
}
