createpostgres:
	docker run --name postgressauth -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgressauth psql

createdb:
	docker exec -it postgressauth createdb --username=root --owner=root go-auth

migrate-addtable:
	migrate create -ext sql -dir DB/migrations add_user_table

migrate-up:
	migrate -path DB/migrations -database "postgresql://root:password@localhost:5433/go-auth?sslmode=disable" -verbose up