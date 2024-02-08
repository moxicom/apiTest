.PHONY: run migration

run:
	go run cmd/main.go

migrate:
	go run cmd/main.go migration
