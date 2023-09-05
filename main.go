package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// menambah tag sql pada struct Employee
type Employees struct {
	ID      int    `sql:"id"`
	Name    string `sql:"name"`
	Age     int    `sql:"age"`
	Address string `sql:"address"`
	Salary  int    `sql:"salary"`
}

func Connect() (*sql.DB, error) {
	dns := "host=localhost user=postgres password=123 dbname=test_db_camp port=5435 sslmode=disable"

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// buat koneksi ke database menggunakan func `Connect`
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// create table employee
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		age INT,
		address VARCHAR(255),
		salary INT
	);`)

	// insert data employee
	_, err = db.Exec(`INSERT INTO employees (name, age, address, salary) VALUES ('John Doe', 30, 'New York', 2000);`)

	// melakukan query untuk mendapatkan semua data dari tabel employee
	row := db.QueryRow("SELECT * FROM employees WHERE id = 1")
	// membuat struct baru untuk menampung data dari tabel employee
	var listEmployee []Employees

	// melakukan looping untuk menampung data dari rows ke struct Employee

	var employees Employees

	// kita tampung setiap baris data ke struct Employee
	err = row.Scan(&employees.ID, &employees.Name, &employees.Age, &employees.Address, &employees.Salary)
	if err != nil {
		panic(err)
	}

	// kemudian kita tambahkan struct Employee ke listEmployee
	listEmployee = append(listEmployee, employees)

	fmt.Println("Data Employee")
	// menampilkan data dari listEmployee
	fmt.Println(listEmployee)

}
