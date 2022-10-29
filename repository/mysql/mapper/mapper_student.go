package mapper

import (
	"fmt"
	"time"
	"try/go-rest/entity"
	"try/go-rest/models"

	"github.com/rocketlaunchr/dbq/v2"
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

func StudentEntityToModel(student *entity.Student) *models.ModelStudent {
	date, _ := time.Parse("2006-01-02", student.GetDob())

	return &models.ModelStudent{
		Nim:          student.GetNim(),
		Name:         student.GetName(),
		Gender:       student.GetGender(),
		Dob:          date,
		Pob:          student.GetPob(),
		Jenjang:      student.GetJenjang(),
		StudyProgram: student.GetStudyProgram(),
		Faculty:      student.GetFaculty(),
	}
}

func StudentEntityToDbqStruct(student *entity.Student) []interface{} {
	dbqStruct := dbq.Struct(StudentEntityToModel(student))
	return dbqStruct
}
