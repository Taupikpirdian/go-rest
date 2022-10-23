package student_handler

import (
	"context"
	_interface "try/go-rest/entity/interface"
)

type StudentHandler struct {
	ctx         context.Context
	repoStudent _interface.InterfaceRepoStudent
}

func NewStudentHandler(ctx context.Context, repoStudent _interface.InterfaceRepoStudent) *StudentHandler {
	return &StudentHandler{
		ctx:         ctx,
		repoStudent: repoStudent,
	}
}
