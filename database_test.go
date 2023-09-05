package belajar-go-database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar-go-database")
	db, err := sql.Open("pq", "host=localhost user=postgres password=123 dbname=test_db_camp port=5435 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// gunakan DB
}
