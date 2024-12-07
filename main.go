package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"karyawan-app-be/api"
	db "karyawan-app-be/db/sqlc"
	"log"
)

func main() {
	// TODO: Load Config
	// TODO: Init DB
	conn, err := sql.Open("postgres", "postgresql://root:12345@localhost:5433/karyawan_app?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start("localhost:8081")

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
