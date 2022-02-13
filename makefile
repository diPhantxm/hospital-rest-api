.PHONY: build
build:
	go build -v -o ./.bin/hospital-rest-api/ ./cmd/hospital-rest-api/

run:
	make build
	go run -v ./cmd/hospital-rest-api/

migrate:
	migrate -path migrations -database "sqlserver://apiserver:1e2w3d4r@192.168.0.179:3306/MSSQLSERVER?database=hospital-rest-api" up

DEFAULT_GOAL: build