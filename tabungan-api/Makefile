createdb:
	docker exec -it postgres-15-test createdb --username=root --owner=root isi_backend_assessment_test

dropdb:
	docker exec -it postgres-15-test dropdb isi_backend_assessment_test

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/isi_backend_assessment_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/isi_backend_assessment_test?sslmode=disable" -verbose down

sqlcwithpowershell:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate

.PHONY: createdb dropdb migrateup migratedown sqlcwithpowershell