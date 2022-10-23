package http_request

type RequestStudent struct {
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
