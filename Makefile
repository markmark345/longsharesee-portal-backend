.PHONY: postgres createdb dropdb miagrateup miagratedown sqlc test start

init:
	go mod tidy

postgres:
	docker run --name air-line-reservation-backend-postgresql-1 -p 5050:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

createdb: 
	docker exec -it air-line-reservation-backend-postgresql-1 createdb --username=root --owner=root db_air_line_reservation

dropdb:
	docker exec -it air-line-reservation-backend-postgresql-1 dropdb db_air_line_reservation

miagrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5050/db_air_line_reservation?sslmode=disable" -verbose up

miagratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5050/db_air_line_reservation?sslmode=disable" -verbose down

test:
	go test ./test/... -coverpkg=./... -coverprofile=./coverage/coverage.out 
	go tool cover -func=./coverage/coverage.out 
	go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
	
start:
	go run cmd/server/main.go
