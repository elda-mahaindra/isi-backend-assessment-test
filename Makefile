postgres:
	docker run --name postgres-15-test -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres-15-test createdb --username=root --owner=root isi_backend_assessment_test

dropdb:
	docker exec -it postgres-15-test dropdb isi_backend_assessment_test

migrateinit:
	migrate create -ext sql -dir db_migration -seq init_schema

migrateup:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/isi_backend_assessment_test?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/isi_backend_assessment_test?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown