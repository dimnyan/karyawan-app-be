package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

//var testDB *sql.DB

func TestMain(m *testing.M) {
	testDB, err := sql.Open("postgres", "postgresql://root:12345@localhost:5433/karyawan_app?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
