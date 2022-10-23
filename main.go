package main

import (
	"context"
	"net/http"
	"try/go-rest/handler/student_handler"
	"try/go-rest/pkg/database/mysql"
	"try/go-rest/pkg/redis"

	repo "try/go-rest/repository/mysql"

	"github.com/gorilla/mux"
)

var (
	mysqlConn   = mysql.InitMysqlDB()
	redisClient = redis.InitRedisClient()
	repoStudent = repo.NewStudnetMysql(mysqlConn, redisClient)
	ctx         = context.Background()
)

func main() {
	r := mux.NewRouter()

	/*
		routes student
	*/
	handlerStudent := student_handler.NewStudentHandler(ctx, repoStudent)
	r.HandleFunc("/create-student", handlerStudent.StoreDataBuku).Methods(http.MethodPost)

}
