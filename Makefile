build:
	@echo " >> building majoo-service binary"
	@go build -v -o majoo-service main.go

run: build
	@./majoo-service

migration-init:
	@go run migrations/*.go init

migration-up:
	@go run migrations/*.go up

migration-down:
	@go run migrations/*.go down

