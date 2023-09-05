package belajar-go-database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar-go-database?parseTime=true")

	db, err := sql.Open("pq", "host=localhost user=postgres password=123 dbname=test_db_camp port=5435 sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // maksimal koneksi yang idle
	db.SetMaxOpenConns(100)                 // maksimal koneksi yang dibuka
	db.SetConnMaxIdleTime(5 * time.Minute)  // koneksi idle selama 5 menit
	db.SetConnMaxLifetime(60 * time.Minute) // koneksi dibuka selama 60 menit

	return db
}
