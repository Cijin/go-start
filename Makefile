dbVersion := 15
dbPort := 5432
dbPassword := password
dbName := dbName
username := root

postgres:
	docker run --name postgres$(dbVersion) -p $(dbPort):$(dbPort) -e POSTGRES_PASSWORD=$(dbPassword) -d postgres:$(dbVersion)-alpine

createdb:
	docker exec -it $(db) createdb --username=$(username) --owner=$(username) $(dbName)

dropdb:
	docker exec -it postgres13 dropdb dbName

migrateup:
	migrate -path db/migration -database "postgresql://$(username):$(password)@localhost:$(dbPort)/$(dbName)" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://$(username):$(password)@localhost:$(dbPort)/$(dbName)" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://$(username):$(password)@localhost:$(dbPort)/$(dbName)" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://$(username):$(password)@localhost:$(dbPort)/$(dbName)" -verbose down 1

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown
