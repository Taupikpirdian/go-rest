package entity

import (
	"errors"
	"time"
)

type Student struct {
	nim           string
	name          string
	gender        string
	dob           time.Time
	pob           string
	jenjang       string
	study_program string
	faculty       string
	created_at    time.Time
	updated_at    time.Time
}

type DTOStudent struct {
	Id           int
	Nim          string
	Name         string
	Gender       string
	Dob          string
	Pob          string
	Jenjang      string
	StudyProgram string
	Faculty      string
	CreatedAt    string
	UpdatedAt    string
}

func NewStudent(dto DTOStudent) (*Student, error) {
	if dto.Nim == "" {
		return nil, errors.New("NIM TIDAK BOLEH KOSONG")
	}

	if dto.Name == "" {
		return nil, errors.New("NAMA TIDAK BOLEH KOSONG")
	}

	if dto.Gender == "" {
		return nil, errors.New("GENDER TIDAK BOLEH KOSONG")
	}

	/*
		#ask, kenapa dto.Dob menggunakan string?
		#answer, supaya mudah validasi ketika data kosong,
		ketika akan masuk ke struct utama, tinggal ubah string to time.
	*/
	if dto.Dob == "" {
		return nil, errors.New("TANGGAL TIDAK BOLEH KOSONG")
	}

	if dto.Pob == "" {
		return nil, errors.New("POB TIDAK BOLEH KOSONG")
	}

	if dto.Jenjang == "" {
		return nil, errors.New("JENJANG TIDAK BOLEH KOSONG")
	}

	if dto.StudyProgram == "" {
		return nil, errors.New("PRODI TIDAK BOLEH KOSONG")
	}

	if dto.Faculty == "" {
		return nil, errors.New("FAKULTAS TIDAK BOLEH KOSONG")
	}

	strDob, _ := time.Parse("2006-01-02", dto.Dob)
	created_at := generateTime()
	updated_at := generateTime()

	if dto.CreatedAt != "" {
		created_at, _ = time.Parse("2006-01-02 15:04:05", dto.CreatedAt)
	}
	if dto.UpdatedAt != "" {
		updated_at, _ = time.Parse("2006-01-02 15:04:05", dto.UpdatedAt)
	}

	student := &Student{
		nim:           dto.Nim,
		name:          dto.Name,
		gender:        dto.Gender,
		dob:           strDob,
		pob:           dto.Pob,
		jenjang:       dto.Jenjang,
		study_program: dto.StudyProgram,
		faculty:       dto.Faculty,
		created_at:    created_at,
		updated_at:    updated_at,
	}

	return student, nil
}

func CollectStudent(dto DTOStudent) (*Student, error) {
	if dto.Nim == "" {
		return nil, errors.New("NIM TIDAK BOLEH KOSONG")
	}

	if dto.Name == "" {
		return nil, errors.New("NAMA TIDAK BOLEH KOSONG")
	}

	if dto.Gender == "" {
		return nil, errors.New("GENDER TIDAK BOLEH KOSONG")
	}

	if dto.Dob == "" {
		return nil, errors.New("TANGGAL TIDAK BOLEH KOSONG")
	}

	if dto.Pob == "" {
		return nil, errors.New("POB TIDAK BOLEH KOSONG")
	}

	if dto.Jenjang == "" {
		return nil, errors.New("JENJANG TIDAK BOLEH KOSONG")
	}

	if dto.StudyProgram == "" {
		return nil, errors.New("PRODI TIDAK BOLEH KOSONG")
	}

	if dto.Faculty == "" {
		return nil, errors.New("FAKULTAS TIDAK BOLEH KOSONG")
	}

	/*
		#bugs: convert jadi data 0001-01-01
	*/
	strDob, _ := time.Parse("2006-01-02", dto.Dob)
	strCreatedAt, _ := time.Parse("2006-01-02 15:04:05", dto.CreatedAt)
	strUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", dto.UpdatedAt)

	student := &Student{
		nim:           dto.Nim,
		name:          dto.Name,
		gender:        dto.Gender,
		dob:           strDob,
		pob:           dto.Pob,
		jenjang:       dto.Jenjang,
		study_program: dto.StudyProgram,
		faculty:       dto.Faculty,
		created_at:    strCreatedAt,
		updated_at:    strUpdatedAt,
	}

	return student, nil
}

/*
	func getter
*/
func (s *Student) GetNim() string {
	return s.nim
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetGender() string {
	return s.gender
}

func (s *Student) GetDob() string {
	return s.dob.Format("2006-01-02")
}

func (s *Student) GetPob() string {
	return s.pob
}

func (s *Student) GetJenjang() string {
	return s.jenjang
}

func (s *Student) GetStudyProgram() string {
	return s.study_program
}

func (s *Student) GetFaculty() string {
	return s.faculty
}

func (s *Student) GetCreatedAt() string {
	return s.created_at.Format("2006-01-02")
}

func (s *Student) GetUpdatedAt() string {
	return s.updated_at.Format("2006-01-02")
}

func generateTime() time.Time {
	data := time.Now()
	return data
}
