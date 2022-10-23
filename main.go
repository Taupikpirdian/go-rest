package main

import (
	"context"
	"fmt"
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
	/*
		DETAIL APLIKASI ADA PADA FILE README.md
	*/
	r := mux.NewRouter()

	// routes student
	handlerStudent := student_handler.NewStudentHandler(ctx, repoStudent)
	r.HandleFunc("/store-student", handlerStudent.StoreDataBuku).Methods(http.MethodPost)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}
