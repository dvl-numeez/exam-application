build:
	@go build -o bin/ExamApplication

run:build
	@./bin/ExamApplication

test:
	@go test -v ./...
