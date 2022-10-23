package mapper

import (
	"try/go-rest/entity"
	"try/go-rest/models"
)

func StudenModelToEntity(model *models.ModelStudent) (*entity.Student, error) {
	student, err := entity.NewStudent(entity.DTOStudent{
		Nim:          model.Nim,
		Name:         model.Name,
		Gender:       model.Gender,
		Dob:          model.Dob,
		Pob:          model.Pob,
		Jenjang:      model.Jenjang,
		StudyProgram: model.StudyProgram,
		Faculty:      model.Faculty,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	})

	if err != nil {
		return nil, err
	}

	return student, nil
}
