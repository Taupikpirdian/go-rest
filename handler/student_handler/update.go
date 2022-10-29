package student_handler

import (
	"encoding/json"
	"net/http"
	"os"
	"try/go-rest/http_request"
	"try/go-rest/http_response"

	"github.com/gorilla/mux"
)

func (s_handler *StudentHandler) UpdateDataStudent(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("api-key")
	if token != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Not Authorization"))
		return
	}

	var (
		req     http_request.RequestStudent
		decoder = json.NewDecoder(r.Body)
		vars    = mux.Vars(r)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	dataStudent, errGet := s_handler.repoStudent.GetStudentByNim(s_handler.ctx, vars["nim"])
	if errGet != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errGet.Error()))
		return
	}

	dataStudent.SetUpdateData(req)

	errUpdate := s_handler.repoStudent.UpdateStudentByNim(s_handler.ctx, dataStudent, vars["nim"])
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errUpdate.Error()))
		return
	}

	resp, errMap := http_response.MapResponseStudent(nil, http.StatusOK, "SUCCESS UPDATE DATA")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMap.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
}
