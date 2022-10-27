package mapper

import (
	"fmt"
	"try/go-rest/entity"
	"try/go-rest/models"
)

func StudenModelToEntity(model *models.ModelStudent) (*entity.Student, error) {
	fmt.Println(model.CreatedAt)
	student, err := entity.NewStudent(entity.DTOStudent{
		Nim:          model.Nim,
		Name:         model.Name,
		Gender:       model.Gender,
		Dob:          model.Dob.Format("2006-01-02"),
		Pob:          model.Pob,
		Jenjang:      model.Jenjang,
		StudyProgram: model.StudyProgram,
		Faculty:      model.Faculty,
		CreatedAt:    model.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    model.UpdatedAt.Format("2006-01-02 15:04:05"),
	})

	if err != nil {
		return nil, err
	}

	return student, nil
}

func DataStudentDbToEntity(dataDTO entity.DTOStudent) (*entity.Student, error) {
	student, err := entity.CollectStudent(dataDTO)
	if err != nil {
		return nil, err
	}

	return student, nil
}
