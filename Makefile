build:
	@go build -o bin/golisten

run: build
	@./bin/golisten