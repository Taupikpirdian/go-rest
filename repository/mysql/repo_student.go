package mysql

import (
	"context"
	"database/sql"
	"try/go-rest/entity"

	"github.com/go-redis/redis/v8"
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

func (b *StudentMysqlInteractor) InsertDataStudent(ctx context.Context, dataBuku *entity.Student) error {
	// var (
	// 	errMysql error
	// )

	// ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	// defer cancel()

	// insertQuery := "INSERT INTO buku (id_pengarang, judul, category, tahun_terbit, kode_buku)" +
	// 	"VALUES(?, ?, ?, ?, ?)"

	// _, errMysql = b.db.Exec(insertQuery, dataBuku.GetDataPengarang().GetId(), dataBuku.GetJudul(), dataBuku.GetCategory(),
	// 	dataBuku.GetTahunTerbit(), dataBuku.GetKodeBuku())

	// if errMysql != nil {
	// 	return errMysql
	// }

	return nil
}
