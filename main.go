package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"try/go-rest/handler/student_handler"
	"try/go-rest/pkg/database/mysql"
	"try/go-rest/pkg/redis"

	repo "try/go-rest/repository/mysql"

	"github.com/gorilla/mux"
)

var (
	_           = setEnvironment()
	mysqlConn   = mysql.InitMysqlDB()
	redisClient = redis.InitRedisClient()
	repoStudent = repo.NewStudnetMysql(mysqlConn, redisClient)
	ctx         = context.Background()
)

func setEnvironment() bool {
	// set env variable using os package
	os.Setenv("APP_NAME", "GO REST STUDENT")
	os.Setenv("APP_ENV", "local")
	os.Setenv("API_KEY", "base64:qIub0xvzVezh+GaRKZQbd3krXge6AnnLghLVvMBbyFM=")

	os.Setenv("DB_CONNECTION", "mysql")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_DATABASE", "go_rest")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "")

	return true
}

func main() {
	/*
		DETAIL APLIKASI ADA PADA FILE README.md
	*/
	// os package
	r := mux.NewRouter()

	// routes student
	handlerStudent := student_handler.NewStudentHandler(ctx, repoStudent)
	r.HandleFunc("/list-student", handlerStudent.IndexDataStudent).Methods(http.MethodGet)
	r.HandleFunc("/store-student", handlerStudent.StoreDataStudent).Methods(http.MethodPost)
	r.HandleFunc("/student/{nim}", handlerStudent.GetStudentByNim).Methods(http.MethodGet)
	r.HandleFunc("/destroy/{nim}", handlerStudent.DestroyDataStudent).Methods(http.MethodDelete)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}
