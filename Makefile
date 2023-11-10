createpostgres:
	docker run --name postgresschat -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgresschat psql

createdb:
	docker exec -it postgresschat createdb --username=root --owner=root go-chat

migrate-addtable:
	migrate create -ext sql -dir DB/migrations add_user_table

migrate-up:
	migrate -path DB/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up