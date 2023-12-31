createdb: 
	docker exec -it postgres createdb --username=root --owner=root simple_bank
	
dropdb: 
	docker exec -it postgres dropdb simple_bank

postgres: 
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
	
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@192.168.215.2:5432/simple_bank?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:secret@192.168.215.2:5432/simple_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...



.PHONY: createdb dropdb postgres migratedown migrateup sqlc test