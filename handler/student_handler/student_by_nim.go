package student_handler

import (
	"net/http"
	"os"
	"try/go-rest/http_response"

	"github.com/gorilla/mux"
)

func (s_handler *StudentHandler) GetStudentByNim(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("api-key")
	if token != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Not Authorization"))
		return
	}

	vars := mux.Vars(r)
	dataBuku, err := s_handler.repoStudent.GetStudentByNim(s_handler.ctx, vars["nim"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, errMap := http_response.MapResponseStudent(dataBuku, 200, "Success")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
