postgres:
	docker run --name dimnyan-psql -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345 -p 5433:5432 -d postgres:alpine

startdb:
	docker container start dimnyan-psql

createdb:
	docker exec -it dimnyan-psql createdb --username=root --owner=root karyawan_app

dropdb:
	docker exec -it dimnyan-psql dropdb karyawan_app

connectdb:
	docker exec -it dimnyan-psql bash
	#psql -U root

migratecreate:
	migrate create -ext sql -dir db/migration/ -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:12345@localhost:5433/karyawan_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:12345@localhost:5433/karyawan_app?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY:
	postgres startdb createdb dropdb migratecreate migrateup migratedown sqlc test server