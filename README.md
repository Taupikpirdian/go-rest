[WHAT_TASK]

[HOW_TO_RUN]
* Installation
  - remove go.mod and go.sum
  - go mod init
  - go mod tidy
  - go mod vendor
  - import db to database (go-rest.sql)
  - go run main.go

[DETAIL_PROJECT]
* DETAIL
  - Ini merupakan aplikasi pendataan mahasiswa/student berprestasi, terdapat 3 table dengan 1 table master dan 2 table transaksi
  - students: master table
  - student_achievements: transaksi table
  - student_achievement_members: transaksi
* TUJUAN
  - Untuk mendata ada berapa prestasi yang dimiliki oleh kampus/lembaga dan siapa saja mahasiswa/studen yang berprestasi tersebut
  - student_achievements: adalah table untuk menampung data prestasi apa saja yang dimiliki kampus/lembaga
  - student_achievement_members: siapa saja yang memiliki prestasi tersebut, relasi one to many dari student_achievements ke student_achievement_members
  - student: data table master dari mahasiswa

* LIST API
  1. Store Data Student
  - curl --location --request POST 'localhost:8080/store-student' \
--header 'Content-Type: application/json' \
--data-raw '{
    "nim": "41037006151017",
    "name": "Taupik Pirdian",
    "gender": "male",
    "dob": "1996-03-19",
    "pob": "Bandung",
    "jenjang": "S1",
    "study_program": "Teknik Informatika",
    "faculty": "Teknik"
}'