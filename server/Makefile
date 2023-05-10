postgres:
	PGPASSWORD=password psql -U postgres -l

database:
	PGPASSWORD=password psql -U postgres -d go-chat -c "SELECT * FROM USERS"

createdb:
	PGPASSWORD=password createdb --username=postgres --owner=postgres go-chat

dropdb:
	PGPASSWORD=password dropdb --username=postgres go-chat

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/go-chat?sslmode=disable" -verbose down

.PHONY: database postgres createdb dropdb migrateup migratedown


prod-postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

prod-postgres:
	docker exec -it postgres15 psql

prod-createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go-chat

prod-dropdb:
	docker exec -it postgres15 dropdb go-chat

prod-migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

prod-migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down

.PHONY: prod-postgresinit prod-postgres prod-createdb prod-dropdb prod-migrateup prod-migratedown