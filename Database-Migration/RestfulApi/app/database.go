package app

import (
	"database/sql"
	"restful-api/helper"
	"time"
)

func NewDB() *sql.DB {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang-database-migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

//migrate
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path RestfulApi/db/migration up

//rollback
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path RestfulApi/db/migration down

//Migrasi ke Versi Tertentu (belakangnya tambahkan angka sesuai dengan versi yang diinginkan)
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path RestfulApi/db/migration up 1

//bikin migrate
//migrate create -ext sql -dir RestfulApi/db/migration create_table_first
//migrate create -ext sql -dir RestfulApi/db/migration create_table_second
//migrate create -ext sql -dir RestfulApi/db/migration create_table_third

//mengetahui versi migrasi
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path RestfulApi/db/migration version

//dirty state menggubah versi (force)
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path RestfulApi/db/migration force <version>