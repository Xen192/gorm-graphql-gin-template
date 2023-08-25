serve:
	go run cmd/app/main.go

gen-graph:
	go run github.com/99designs/gqlgen generate

gen-model:
	go run cmd/model/generate.go

migrate:
	go run cmd/migrate/main.go gen apply

migrate-gen:
	go run cmd/migrate/main.go gen

migrate-apply:
	go run cmd/migrate/main.go apply

migrate-dry:
	go run cmd/migrate/main.go apply dry

migrate-new:
	atlas migrate new --dir "file://db/migration"

migrate-hash:
	atlas migrate hash --dir "file://db/migration"

docker-db:
	docker container rm -f mygpt-db && docker-compose -f docker-compose.local.yml run --rm -d --service-ports --name mygpt-db postgres
