package student_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"try/go-rest/entity"
	"try/go-rest/http_request"
)

func (s_handler *StudentHandler) StoreDataStudent(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestStudent
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	/*
		#minor_bugs: jam created at masih 00:00:00
	*/
	student, err := entity.NewStudent(entity.DTOStudent{
		Nim:          req.Nim,
		Name:         req.Name,
		Gender:       req.Gender,
		Dob:          req.Dob,
		Pob:          req.Pob,
		Jenjang:      req.Jenjang,
		StudyProgram: req.StudyProgram,
		Faculty:      req.Faculty,
	})

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error build data"))
		return
	}

	/*
		Cek data student by nim, jika sudah ada jangan save lagi
	*/
	checkDataByNim, errDataByNim := s_handler.repoStudent.CheckDataStudentByNim(s_handler.ctx, req.Nim)
	if errDataByNim != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errDataByNim.Error()))
		return
	}

	if !checkDataByNim {
		fmt.Fprintf(w, "DATA DENGAN NIM: "+req.Nim+", SUDAH ADA")
		return
	}

	errInsert := s_handler.repoStudent.InsertDataStudent(s_handler.ctx, student)
	if errInsert != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errInsert.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "SUCCESS INSERT DATA")
}
