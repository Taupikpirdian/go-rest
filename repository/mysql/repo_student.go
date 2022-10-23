package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"try/go-rest/entity"
	"try/go-rest/models"

	"github.com/go-redis/redis/v8"
	"github.com/rocketlaunchr/dbq/v2"
)

type StudentMysqlInteractor struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewStudnetMysql(db *sql.DB, redisClient *redis.Client) *StudentMysqlInteractor {
	return &StudentMysqlInteractor{
		db:          db,
		redisClient: redisClient,
	}
}

func (b *StudentMysqlInteractor) InsertDataStudent(ctx context.Context, dataStudent *entity.Student) error {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO students (nim, name, gender, dob, pob, jenjang, study_program, faculty, created_at, updated_at)" +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, errMysql = b.db.Exec(insertQuery, dataStudent.GetNim(), dataStudent.GetName(), dataStudent.GetGender(),
		dataStudent.GetDob(), dataStudent.GetPob(), dataStudent.GetJenjang(), dataStudent.GetStudyProgram(), dataStudent.GetFaculty(), dataStudent.GetCreatedAt(), dataStudent.GetUpdatedAt())

	if errMysql != nil {
		return errMysql
	}

	return nil
}

func (s *StudentMysqlInteractor) CheckDataStudentByNim(ctx context.Context, nim string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	queryBuku := fmt.Sprintf("SELECT * FROM %s WHERE nim = ?", models.GetStudentTableName())

	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.ModelStudent{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	resultStudent, err := dbq.Q(ctx, s.db, queryBuku, opts, nim)

	if err != nil {
		return false, err
	}

	if resultStudent == nil {
		return true, nil
	} else {
		return false, errors.New("DATA STUDENT SUDAH ADA DALAM DATABASE")
	}
}
