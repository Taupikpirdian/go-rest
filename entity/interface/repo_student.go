package _interface

import (
	"context"
	"try/go-rest/entity"
)

type InterfaceRepoStudent interface {
	InsertDataStudent(ctx context.Context, dataStudent *entity.Student) error
	CheckDataStudentByNim(ctx context.Context, nim string) (bool, error)
	ListDataStudent(ctx context.Context) ([]*entity.Student, error)
	GetStudentByNim(ctx context.Context, nim string) (*entity.Student, error)
	UpdateStudentByNim(ctx context.Context, dataBuku *entity.Student, kodeBuku string) error
	DeleteDataStudentByNim(ctx context.Context, nim string) error
}
