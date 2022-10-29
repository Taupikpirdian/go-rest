package student_handler

import (
	"net/http"
	"os"
	"try/go-rest/http_response"

	"github.com/gorilla/mux"
)

func (s_handler *StudentHandler) DestroyDataStudent(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("api-key")
	if token != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Not Authorization"))
		return
	}

	var vars = mux.Vars(r)

	dataStudent, errGet := s_handler.repoStudent.GetStudentByNim(s_handler.ctx, vars["nim"])
	if errGet != nil {
		respErr, _ := http_response.MapResponseStudent(nil, http.StatusInternalServerError, errGet.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write(respErr)
		return
	}

	errDelete := s_handler.repoStudent.DeleteDataStudentByNim(s_handler.ctx, dataStudent.GetNim())
	if errDelete != nil {
		respErr, _ := http_response.MapResponseStudent(nil, http.StatusInternalServerError, errDelete.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}

	resp, errMap := http_response.MapResponseStudent(nil, http.StatusOK, "SUCCESS DELETE DATA")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMap.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
}
