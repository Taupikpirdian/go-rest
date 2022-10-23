package models

type ModelStudent struct {
	Nim          string `json:"nim"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	Dob          string `json:"dob"`
	Pob          string `json:"pob"`
	Jenjang      string `json:"jenjang"`
	StudyProgram string `json:"study_program"`
	Faculty      string `json:"faculty"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func GetStudentTableName() string {
	return "students"
}

func GetStudentFillableTable() []string {
	return []string{
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
