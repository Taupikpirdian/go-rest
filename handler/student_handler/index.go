package student_handler

import (
	"net/http"
	"try/go-rest/http_response"
)

func (s_handler *StudentHandler) IndexDataStudent(w http.ResponseWriter, r *http.Request) {
	// #ask, kenapa s_handler bisa akses repoStudent ?
	listStudent, err := s_handler.repoStudent.ListDataStudent(s_handler.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	response, errMap := http_response.MapResponseListStudent(listStudent, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}

	w.WriteHeader(200)
	w.Write(response)
}
