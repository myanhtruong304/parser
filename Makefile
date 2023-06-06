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

seed-db:
	migrate -path db/seed -database "postgresql://root:root@localhost:5433/pay68?sslmode=disable" -verbose up

sqlc:
	sqlc generate

server:
	go run ./cmd/main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/pay68/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown seed-db sqlc test server mock