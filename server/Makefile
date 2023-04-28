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