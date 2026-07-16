build:
	@go build -o ./bin/taskarena ./cmd/taskarena
test:
	@go test -v ./...
run: 	build
	@./bin/taskarena
