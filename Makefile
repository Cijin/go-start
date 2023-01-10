dbVersion := 14
dbPort := 5432
dbPassword := password
dbName := projectDb
dbUser := root
dockerImage := postgres14

# ------------------------- Database
createdb:
	docker start --name $(dockerImage) -p $(dbPort):$(dbPort) -e POSTGRES_DB=$(dbName) -e POSTGRES_USER=$(dbUser) -e POSTGRES_PASSWORD=$(dbPassword) -d postgres:14-alpine

startdb:
	docker start $(dockerImage) 

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

# ------------------------- Start Server
server:
	go run main.go

# ------------------------- Test & Mocks
test: 
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/cijin/go-start/db/sqlc Store

.PHONY: createdb startdb psql createmigration migrateup migrateup1 migratedown migratedown1 sqlc test mock
