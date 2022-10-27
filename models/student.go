package models

import "time"

type ModelStudent struct {
	Id           string    `dbq:"id"`
	Nim          string    `dbq:"nim"`
	Name         string    `dbq:"name"`
	Gender       string    `dbq:"gender"`
	Dob          time.Time `dbq:"dob"`
	Pob          string    `dbq:"pob"`
	Jenjang      string    `dbq:"jenjang"`
	StudyProgram string    `dbq:"study_program"`
	Faculty      string    `dbq:"faculty"`
	CreatedAt    time.Time `dbq:"created_at"`
	UpdatedAt    time.Time `dbq:"updated_at"`
}

func GetStudentTableName() string {
	return "students"
}

func GetStudentFillableTable() []string {
	return []string{
		"id",
		"nim",
		"name",
		"gender",
		"dob",
		"pob",
		"jenjang",
		"study_program",
		"faculty",
		"created_at",
		"updated_at",
	}
}
