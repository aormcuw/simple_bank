postgres:
	docker run --name postgres16 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pororo -d postgres:16.3
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres16 dropdb --username=root --owner=root simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:pororo@localhost:5433/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:pororo@localhost:5433/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb  -destination db/mock/store.go simplebank/db/sqlc Store

.PHONY: createdb dropdb postgres migratedown migrateup sqlc test server mock