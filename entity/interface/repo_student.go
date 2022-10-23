package _interface

import (
	"context"
	"try/go-rest/entity"
)

type InterfaceRepoStudent interface {
	InsertDataStudent(ctx context.Context, dataStudent *entity.Student) error
}
