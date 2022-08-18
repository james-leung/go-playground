runpostgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=postgres --owner=postgres todo_list

dropdb:
	docker exec -it postgres14 dropdb --username=postgres todo_list

migrateup:
	migrate -path db/migration -database "postgres://postgres:admin@localhost:5432/todo_list?sslmode=disable" -verbose up
	docker exec -it postgres14 dropdb --username=postgres todo_list

migratedown:
	migrate -path db/migration -database "postgres://postgres:admin@localhost:5432/todo_list?sslmode=disable" -verbose down

sqlc_generate:
	sqlc generate

.PHONY: runpostgres createdb dropdb migrateup migratedown sqlc_generate