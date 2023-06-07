postgres:
	docker run --name pay68 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=roor -d postgres:15.1-alpine

createdb:
	docker exec -it pay68 createdb --username=root --owner=root pay68

dropdb:
	docker exec -it pay68 dropdb --username=root --owner=root pay68

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5433/pay68?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5433/pay68?sslmode=disable" -verbose down

seed:
	go run ./cmd/db/main.go
sqlc:
	sqlc generate

server:
	go run ./cmd/api/main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/parser/db/sqlc Store

proto:
	rm -f pb/*.go
	$ protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown seed-db sqlc test server mock proto