dbVersion := 14
dbPort := 5432
dbPassword := password
dbName := projectDb
dbUser := root

# ------------------------- Database
startdb:
	docker run --name postgres14 -p $(dbPort):$(dbPort) -e POSTGRES_DB=$(dbName) -e POSTGRES_USER=$(dbUser) -e POSTGRES_PASSWORD=$(dbPassword) -d postgres:14-alpine

psql:
	docker exec -it postgres14 psql -U $(dbUser) -d $(dbName)

# ------------------------- Migration
createmigration:
	migrate create -ext sql -dir db/migration -seq # pass seq here

migrateup:
	migrate -path db/migration -database "postgresql://$(dbUser):$(dbPassword)@localhost:$(dbPort)/$(dbName)?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://$(dbUser):$(dbPassword)@localhost:$(dbPort)/$(dbName)?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://$(dbUser):$(dbPassword)@localhost:$(dbPort)/$(dbName)?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://$(dbUser):$(dbPassword)@localhost:$(dbPort)/$(dbName)?sslmode=disable" -verbose down 1

# ------------------------- Sqlc
sqlc:
	sqlc generate

# ------------------------- Test & Coverage
test: 
	go test -v -cover ./...

.PHONY: startdb psql createmigration migrateup migrateup1 migratedown migratedown1 sqlc test
