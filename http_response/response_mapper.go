package http_response

import (
	"encoding/json"
	"try/go-rest/entity"
)

type Status struct {
	Code    int
	Message string
}

type ResponseStudentJson struct {
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

type CustomReponseCollection struct {
	Status *Status
	Data   []*ResponseStudentJson
}

func MapResponseListStudent(dataStudent []*entity.Student, code int, message string) ([]byte, error) {
	listResp := make([]*ResponseStudentJson, 0)
	for _, data := range dataStudent {
		resp := &ResponseStudentJson{
			Nim:          data.GetNim(),
			Name:         data.GetName(),
			Gender:       data.GetGender(),
			Dob:          data.GetDob(),
			Pob:          data.GetPob(),
			Jenjang:      data.GetJenjang(),
			StudyProgram: data.GetStudyProgram(),
			Faculty:      data.GetFaculty(),
			CreatedAt:    data.GetCreatedAt(),
			UpdatedAt:    data.GetUpdatedAt(),
		}

		listResp = append(listResp, resp)
	}

	httpResponse := &CustomReponseCollection{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: listResp,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
