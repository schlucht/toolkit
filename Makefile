build:
	@clear
	@go build -o bin/gobank

run: build
	@./bin/gobank

test:
	@go test -v ./...

docker-run:
	@docker run --name d-postgres -e POSTGRES_PASSWORD=schlucht -p 5432:5432 -d postgres